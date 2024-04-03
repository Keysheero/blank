package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gostart/internal/config"
	user "gostart/internal/models"
	"net/http"
	"time"
)

type JWTData struct {
	sub string
	exp int
	iat int
}

func GenerateJWTToken(cnfg *config.Config, user user.User, w http.ResponseWriter) string {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(cnfg.Auth.Secret)

	return tokenString

}

func RetriveJwtToken(r *http.Request, secret string) (JWTData, error) {
	cookie, _ := r.Cookie("access_token")
	var data JWTData
	if cookie.Value == "" {
		return data, errors.New("cookie not found")
	} else {
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			return data, err
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			jsonClaims, err := json.Marshal(claims)
			if err != nil {
				return data, err
			}
			err = json.Unmarshal(jsonClaims, &data)
			if err != nil {
				return data, err
			}
		} else {
			return data, fmt.Errorf("invalid token")
		}

	}
	return data, nil

}
