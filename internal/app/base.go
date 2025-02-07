package app

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	usecaseinterfaces "gostart/internal/application/usecases/interfaces"
	user2 "gostart/internal/application/usecases/user"
	serviceinterfaces "gostart/internal/domain/services/interfaces"
	"gostart/internal/domain/services/user"
	"gostart/internal/infrastructure/config"
	"gostart/internal/infrastructure/database/repository/storage/postgres"
	repositoryinterfaces "gostart/internal/infrastructure/database/repository/storage/postgres/interfaces"
	user3 "gostart/internal/infrastructure/database/repository/storage/postgres/user"
	"gostart/internal/infrastructure/logger"
	validatorinterfaces "gostart/internal/infrastructure/validator"
	"gostart/internal/infrastructure/validator/interfaces"
	handlerinterfaces "gostart/internal/transport/http-server/handlers/interfaces"
	base "gostart/internal/transport/http-server/handlers/user"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	Config *config.Config
	Logger logger.Logger

	UserRepository repositoryinterfaces.UserRepository
	UserService    serviceinterfaces.UserService
	UserUsecase    usecaseinterfaces.UserUsecase
	UserValidator  validator.UserValidator
	UserHandler    handlerinterfaces.UserHandler

	Router chi.Router
	Server *http.Server
}

func InitializeApplication() *Application {
	return &Application{}
}

func (app *Application) Run() {
	if err := app.InitializeConfig(); err != nil {
		app.Logger.Error("Config initialize error", "ERROR", err)
	}
	app.InitializeLogger()
	if err := app.InitializeUserRepository(); err != nil {
		app.Logger.Error("Logger initialize error", "ERROR", err)
	}
	app.InitializeUserService()
	app.InitializeUserUsecase()
	app.InitializeUserHandler()
	app.InitializeRouter()

	app.RegisterPostHandlers()
	app.InitializeServer()
	go func() {
		app.Logger.Info("Server Launching", "SERVER_ADDRESS", app.Config.HttpServer.Address)

		err := app.Server.ListenAndServe()

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.Logger.Error("Server startup error", "ERROR", err)
		}

	}()
	app.GracefulShutdown()

}

func (app *Application) InitializeConfig() error {
	cnfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	app.Config = cnfg
	return nil
}

func (app *Application) InitializeLogger() {
	app.Logger = logger.NewSlogger()
}

func (app *Application) InitializeUserRepository() error {
	client, err := postgres.NewClient(context.Background(), app.Config.Database, app.Logger)
	if err != nil {
		app.Logger.Error("Database client initialize error", "ERROR", err)
		return err
	}

	app.UserRepository = user3.NewUserRepository(client, app.Logger, app.Config)
	return nil
}

func (app *Application) InitializeUserService() {
	app.UserService = user.NewUserService(app.UserRepository, app.Logger)
}

func (app *Application) InitializeUserUsecase() {
	app.UserUsecase = user2.NewUserUsecase(app.UserService, app.Logger)
}

func (app *Application) InitializeUserValidator() {
	app.UserValidator = validatorinterfaces.NewUserValidator()
}

func (app *Application) InitializeUserHandler() {
	app.UserHandler = base.NewUserHandler(app.UserUsecase, app.Logger, app.UserValidator)
}

func (app *Application) InitializeRouter() {
	app.Router = chi.NewRouter()
	app.Router.Get("/swagger/*", httpSwagger.WrapHandler)
}

func (app *Application) RegisterPostHandlers() {

}

func (app *Application) InitializeServer() {
	server := &http.Server{
		Addr:         app.Config.HttpServer.Address,
		Handler:      app.Router,
		ReadTimeout:  time.Duration(app.Config.HttpServer.Timeout) * time.Second,
		WriteTimeout: time.Duration(app.Config.HttpServer.Timeout) * time.Second,
		IdleTimeout:  time.Duration(app.Config.HttpServer.IdleTimeout) * time.Second,
	}
	app.Server = server

}

func (app *Application) GracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	app.Logger.Info("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := app.Server.Shutdown(ctx); err != nil {
		app.Logger.Error("Server forced to shutdown", "ERROR", err)
	}

	app.Logger.Info("Server exiting")
}
