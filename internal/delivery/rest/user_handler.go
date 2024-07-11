package rest

import (
	"encoding/json"
	"fmt"
	"golang-api-restaurant/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) RegisterUser(c echo.Context) error {
	var request model.RegisterRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("gor error: %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	userData, err := h.restoUsecase.RegisterUser(request)
	if err != nil {
		fmt.Printf("gor error: %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": userData,
	})

}


func (h *handler) Login(c echo.Context) error {
	var request model.LoginRequest
    err := json.NewDecoder(c.Request().Body).Decode(&request)
    if err!=nil{
        fmt.Printf("gor error: %s\n", err.Error())
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "error": err.Error(),
        })
    }
    sessionData, err:=h.restoUsecase.Login(request)
    if err!=nil{
        fmt.Printf("gor error: %s\n", err.Error())
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "error": err.Error(),
        })
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "data": sessionData,
    })
}