package main

import (
	"golang-api-restaurant/internal/database"
	"golang-api-restaurant/internal/delivery/rest"
	mRepo "golang-api-restaurant/internal/respository/menu"
	rUsecase "golang-api-restaurant/internal/usecase/resto"

	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=postgres dbname=go_resto_app sslmode=disable"
)

func main() {

	e := echo.New()

	db := database.GetDB(dbAddress)

	menuRepo:=mRepo.GetRepository(db)
	restoUsecase:= rUsecase.GetUsecase(menuRepo)
	h:=rest.NewHandler(restoUsecase)

	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":14045"))

}
