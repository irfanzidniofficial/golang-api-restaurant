package rest

import (
	"errors"
	"golang-api-restaurant/internal/model"
	"net/http"
	"strings"
)

func GetSessionData(r *http.Request) (model.UserSession, error) {
	authString := r.Header.Get("Authorization")
	splitString := strings.Split(authString, "Bearer ")
	if len(splitString) != 2 {
		return model.UserSession{}, errors.New("unauthorized")
	}
	accessString := splitString[1]
	return model.UserSession{
		JWTToken: accessString,
	}, nil
}
