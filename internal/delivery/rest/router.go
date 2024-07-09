package rest

import "github.com/labstack/echo/v4"

func LoadRoutes(e *echo.Echo, handler *handler) {
	e.GET("/menu", handler.GetMenuList)
	e.POST("/order", handler.Order)
	e.GET("/order/:orderID", handler.GetOrderInfo)
}
