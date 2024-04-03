package user

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"gostart/internal/config"
	"gostart/internal/database/postgres"
	"gostart/internal/http-server/handlers"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	handlers.Handler
}

func (uh *UserHandler) UserGetHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id, _ := uuid.Parse(r.URL.Query().Get("id"))

	currentUser, _ := uh.Repo.GetByID(context.Background(), id)
	err := json.NewEncoder(w).Encode(map[string]string{"status": "successful", "user": currentUser.Name})
	if err != nil {
		uh.Logger.Warn("Problem with responding user by id", "ERROR", err)
	}

}

func NewUserHandler(repo *postgres.UserRepository, logger *slog.Logger, config *config.Config) *UserHandler {
	return &UserHandler{Handler: handlers.Handler{
		Repo:   repo,
		Logger: logger,
		Config: config,
	}}
}
