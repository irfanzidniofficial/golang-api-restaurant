package main

import (
	"crypto/rand"
	"crypto/rsa"
	"golang-api-restaurant/internal/database"
	"golang-api-restaurant/internal/delivery/rest"
	"golang-api-restaurant/internal/logger"
	mRepo "golang-api-restaurant/internal/respository/menu"
	oRepo "golang-api-restaurant/internal/respository/order"
	uRepo "golang-api-restaurant/internal/respository/user"
	rUsecase "golang-api-restaurant/internal/usecase/resto"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=postgres dbname=go_resto_app sslmode=disable"
)

func main() {
	logger.Init()

	e := echo.New()

	db := database.GetDB(dbAddress)
	secret := "AES256Key-32Characters1234567890"
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32, signKey, 60*time.Second)
	if err != nil {
		panic(err)
	}
	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo, userRepo)
	h := rest.NewHandler(restoUsecase)
	rest.LoadMiddlewares(e)

	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":14045"))

}
