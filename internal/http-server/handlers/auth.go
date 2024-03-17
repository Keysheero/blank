package handlers

import (
	"context"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"gostart/internal/schemas"
	"net/http"
)

type UserHandler struct {
	Handler
}

func (uh *UserHandler) UserRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var data schemas.UserRegistration
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		uh.Logger.Warn("Some Problem with parsing json", "ERROR", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "Ошибка разбора JSON"}); err != nil {
			uh.Logger.Warn("Problem with response data to user", err)
		}
		return
	}

	if err := uh.Repo.Create(context.Background(), data); err != nil {
		uh.Logger.Warn("Problems with create function", "DBERROR", err)
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "Ошибка создания пользователя"}); err != nil {
			uh.Logger.Warn("Problem with response data to user", err)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Пользователь успешно создан"}); err != nil {
		uh.Logger.Warn("Problem with response data to user", err)
	}
}

func (uh *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data schemas.UserLogin

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		uh.Logger.Warn("Problem with decoding(marshaling) to json")
	}
	uh.Logger.Debug("Receiving data from marshalling json", "THE DATA:", data)
}

func ValidatePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
