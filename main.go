package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"gostart/internal/config"
	"gostart/internal/database/postgres"
	"gostart/internal/http-server/handlers/auth"
	"gostart/internal/http-server/handlers/user"
	router2 "gostart/internal/http-server/router"
	logger2 "gostart/internal/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()

	cnfg := config.LoadConfig()

	logger := logger2.SetupLogger(cnfg.Env)

	logger.Info("Application start-up")
	client, _ := postgres.NewClient(ctx, cnfg.Database, logger)
	// USER REPO INIT
	repo := postgres.NewUserRepository(client, logger, cnfg)
	// USER HANDLERS INIT
	handler := user.NewUserHandler(repo, logger, cnfg)
	authHandler := auth.NewAuthHandler(repo, logger, cnfg)

	router := chi.NewRouter()
	router2.SetupMiddlewares(router, logger)
	router2.SetupUserHandlers(handler, router)
	router2.SetupAuthHandler(authHandler, router)

	// server startup
	server := &http.Server{
		Addr:         cnfg.HttpServer.Address,
		Handler:      router,
		ReadTimeout:  time.Duration(cnfg.HttpServer.Timeout),
		WriteTimeout: time.Duration(cnfg.HttpServer.Timeout),
		IdleTimeout:  time.Duration(cnfg.HttpServer.IdleTimeout),
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := server.ListenAndServe(); err != nil && nil != http.ErrServerClosed {
			logger.Warn("Problems while running server", "ERROR", err)
		}
	}()
	<-done
	logger.Info("Shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cnfg.HttpServer.Timeout))

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Warn("Problems while trying to shut down the server", "ERROR", err)
	} else {
		logger.Info("Server shut down Gracefully")
	}

}
