package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
		},
		{
			Name:      "Ayam Rica Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}

func main() {
	e := echo.New()
	e.GET("/menu/food", getFoodMenu)

	e.Logger.Fatal(e.Start(":14045"))

}
