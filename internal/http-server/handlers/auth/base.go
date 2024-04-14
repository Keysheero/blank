package auth

import (
	"context"
	"encoding/json"
	"gostart/internal/config"
	"gostart/internal/database/postgres"
	"gostart/internal/http-server/handlers"
	schemas2 "gostart/internal/http-server/schemas"
	"log/slog"
	"net/http"
)

type AuthHandler struct {
	handlers.Handler
}

func (uh *AuthHandler) UserRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var data schemas2.UserRegistration
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		uh.Logger.Warn("Some Problem with parsing json", "ERROR", err)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "Ошибка разбора JSON"}); err != nil {
			uh.Logger.Warn("Problem with response data to user", err)
		}
		return
	}

	if err := uh.Repo.NewUser(context.Background(), data); err != nil {
		uh.Logger.Warn("Problems with create function", "DBERROR", err)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "Ошибка создания пользователя"}); err != nil {
			uh.Logger.Warn("Problem with response data to user", err)
			return
		}
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Пользователь успешно создан"}); err != nil {
		uh.Logger.Warn("Problem with responding data to user", err)
	}
}

//func (uh *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
//	var data schemas2.UserLogin
//
//	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
//		uh.Logger.Warn("Problem with decoding json")
//		http.Error(w, "Ошибка разбора json", http.StatusBadRequest)
//		return
//	}
//	user, ok := uh.Repo.ValidateUser(context.Background(), data.Username, data.Password)
//	if !ok {
//		http.Error(w, "Wrong data", http.StatusUnauthorized)
//		return
//	}
//	token := jwt.GenerateJWTToken(uh.Config, user, w)
//
//	http.SetCookie(w, &http.Cookie{
//		Name:     "access_token",
//		Value:    token,
//		Expires:  time.Now().Add(time.Hour * 1), // Срок жизни куки 1 час
//		Path:     "/",
//		HttpOnly: true,
//		SameSite: http.SameSiteStrictMode,
//	})
//	w.Header().Set("Content-Type", "application/json")
//	_ = json.NewEncoder(w).Encode(schemas2.LoginResponse{AccessToken: token})
//}

func (uh *AuthHandler) ProtectedResource(w http.ResponseWriter, r *http.Request) {

}

func NewAuthHandler(repo *postgres.UserRepository, logger *slog.Logger, config *config.Config) *AuthHandler {
	return &AuthHandler{Handler: handlers.Handler{
		Repo:   repo,
		Logger: logger,
		Config: config,
	}}
}
