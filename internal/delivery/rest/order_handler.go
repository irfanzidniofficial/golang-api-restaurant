package rest

import (
	"encoding/json"
	"fmt"
	"golang-api-restaurant/internal/model"
	"golang-api-restaurant/internal/model/constant"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) Order(c echo.Context) error {
	var request model.OrderMenuRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("gor error: %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	userID := c.Request().Context().Value(constant.AuthContextKey).(string)
	request.UserID = userID

	orderData, err := h.restoUsecase.Order(request)
	if err != nil {
		fmt.Printf("gor error: %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": orderData,
	})
}

func (h *handler) GetOrderInfo(c echo.Context) error {
	orderID := c.Param("orderID")
	userID := c.Request().Context().Value(constant.AuthContextKey).(string)

	orderData, err := h.restoUsecase.GetOrderInfo(model.GetOrderInfoRequest{
		UserID:  userID,
		OrderID: orderID,
	})
	if err != nil {
		fmt.Printf("gor error: %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": orderData,
	})
}
