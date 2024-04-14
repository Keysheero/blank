package user

import (
	"context"
	"encoding/json"
	"fmt"
	"gostart/internal/config"
	"gostart/internal/database/postgres"
	"gostart/internal/http-server/handlers"
	"gostart/internal/http-server/schemas"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	handlers.Handler
}

//func (uh *UserHandler) UserGetHandler(w http.ResponseWriter, r *http.Request) {
//	defer r.Body.Close()
//	id, _ := uuid.Parse(r.URL.Query().Get("id"))
//
//	currentUser, _ := uh.Repo.GetByID(context.Background(), id)
//	err := json.NewEncoder(w).Encode(map[string]string{"status": "successful", "user": currentUser.Name})
//	if err != nil {
//		uh.Logger.Warn("Problem with responding user by id", "ERROR", err)
//	}
//
//}

func (uh *UserHandler) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {

	student, err := uh.Repo.GetUserByEmail(context.Background(), "huyl1o@gmail.com")
	if err != nil {
		uh.Logger.Warn("HELLO-WORLD HANDLER HAS FUCKED UP")
	}
	if _, err := fmt.Fprintf(w, "hello world"+student.Email); err != nil {
		uh.Logger.Warn("PROblem heree", err)
	}

}

func (uh *UserHandler) PostTestHandler(w http.ResponseWriter, r *http.Request) {
	var data schemas.UserRegistration
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		uh.Logger.Debug("POST HANDLER DECODING PROBLEM", "ERROR", err)
	}
	//if err := uh.Repo.NewUser(context.Background(), data); err != nil {
	//	uh.Logger.Info("POST-HANDLER special problem was fucked up well too", err)
	//}
	uh.Logger.Debug("POST HANDLER HAS DECODED JSON SUCCESSFULLY")
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "ZAEBIS"}); err != nil {
		uh.Logger.Warn("POST HANDLER PROBLEM TO RESPONSE DATA", err)
	}
}

func NewUserHandler(repo *postgres.UserRepository, logger *slog.Logger, config *config.Config) *UserHandler {
	return &UserHandler{Handler: handlers.Handler{
		Repo:   repo,
		Logger: logger,
		Config: config,
	}}
}
