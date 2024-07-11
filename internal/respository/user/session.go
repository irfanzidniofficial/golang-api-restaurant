package user

import (
	"errors"
	"golang-api-restaurant/internal/model"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
}

func (ur *userRepo) CreateUserSession(userID string) (model.UserSession, error) {
	accessToken, err := ur.generateAccessToken(userID)
	if err != nil {
		return model.UserSession{}, err
	}
	return model.UserSession{
		JWTToken: accessToken,
	}, nil
}

func (ur *userRepo) generateAccessToken(userID string) (string, error) {
	accessTokenExp := time.Now().Add(ur.accessExp).Unix()
	accessClaim := Claims{
		jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: accessTokenExp,
		},
	}
	accesJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), accessClaim)

	return accesJwt.SignedString(ur.signKey)
}

func (ur *userRepo) CheckSession(data model.UserSession) (userID string, err error) {
	accessToken, err := jwt.ParseWithClaims(data.JWTToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return &ur.signKey.PublicKey, nil
	})
	if err != nil {
		return "", err
	}
	accessTokenClaims, ok := accessToken.Claims.(*Claims)
	if !ok {
		return "", errors.New("unauthorized access")
	}

	if accessToken.Valid {
		return accessTokenClaims.Subject, nil
	}
	return "", errors.New("unauthorized access")
}
