package handlers

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func (uh *UserHandler) UserGetHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id, _ := uuid.Parse(r.URL.Query().Get("id"))

	currentUser, _ := uh.Repo.GetByID(context.Background(), id)
	err := json.NewEncoder(w).Encode(map[string]string{"status": "successful", "user": currentUser.Name})
	if err != nil {
		uh.Logger.Warn("Problem with responding user by id", "ERROR", err)
	}

}
