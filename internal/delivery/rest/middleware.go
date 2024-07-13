package rest

import (
	"context"
	"golang-api-restaurant/internal/model/constant"
	"golang-api-restaurant/internal/tracking"
	"golang-api-restaurant/internal/usecase/resto"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoadMiddlewares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://restoku.com"},
	}))
}

func GetAuthMiddleware(restoUsecase resto.Usecase) *authMiddleware {
	return &authMiddleware{
		restoUsecase: restoUsecase,
	}
}

type authMiddleware struct {
	restoUsecase resto.Usecase
}

func (am *authMiddleware) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracking.CreateSpan(c.Request().Context(), "CheckAuth")
		defer span.End()
		sessionData, err := GetSessionData(c.Request())
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  err.Error(),
				Internal: err,
			}

		}
		userID, err := am.restoUsecase.CheckSession(ctx, sessionData)
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  err.Error(),
				Internal: err,
			}
		}
		authContext := context.WithValue(c.Request().Context(), constant.AuthContextKey, userID)
		c.SetRequest(c.Request().WithContext(authContext))
		if err := next(c); err != nil {
			return err
		}
		return nil
	}
}
