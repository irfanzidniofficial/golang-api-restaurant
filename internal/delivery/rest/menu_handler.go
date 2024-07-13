package rest

import (
	"fmt"
	"golang-api-restaurant/internal/tracing"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetMenuList(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "GetMenuList")
	defer span.End()

	menuType := c.FormValue("menu_type")

	menuData, err := h.restoUsecase.GetMenuList(ctx, menuType)
	if err != nil {
		fmt.Printf("got error: %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}
