package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gostart/internal/http-server/handlers/auth"
	"gostart/internal/http-server/handlers/user"
	mwLogger "gostart/internal/http-server/middleware"
	"log/slog"
)

func SetupMiddlewares(router chi.Router, logger *slog.Logger) {
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(logger))
	router.Use(middleware.URLFormat)
	router.Use(middleware.Recoverer)
}

func SetupUserHandlers(h *user.UserHandler, router chi.Router) {
	router.Get("/user/get", h.UserGetHandler)
}

func SetupAuthHandler(h *auth.AuthHandler, router chi.Router) {
	router.Post("/auth/new/", h.UserRegistrationHandler)
	router.Post("/auth/login/", h.LoginHandler)
}
