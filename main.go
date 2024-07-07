package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=postgres dbname=go_resto_app sslmode=disable"
)

type MenuType string

const (
	MenuTypeFood  = "food"
	MenuTypeDrink = "drink"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      MenuType
}

func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}
	var menuData []MenuItem

	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

func seedDB() {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Ayam Rica Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      MenuTypeFood,
		},
	}

	drinksMenu := []MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     4000,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Es Teh Manis",
			OrderCode: "es_teh_manis",
			Price:     5000,
			Type:      MenuTypeDrink,
		},
	}

	fmt.Println(foodMenu)
	fmt.Println(drinksMenu)

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	err = db.AutoMigrate(&MenuItem{})
	if err != nil {
		panic(err)
	}
	var count int64
	db.Model(&MenuItem{}).Count(&count)
	if count == 0 {
		// Insert seed data
		if err := db.Create(&foodMenu).Error; err != nil {
			panic(err)
		}
		if err := db.Create(&drinksMenu).Error; err != nil {
			panic(err)
		}
		fmt.Println("Seed data inserted successfully")
	} else {
		fmt.Println("Seed data already exists")
	}

	// Insert seed data
	// if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
	// 	db.Create(&foodMenu)
	// 	db.Create(&drinksMenu)
	// }

}

func main() {
	seedDB()
	e := echo.New()
	e.GET("/menu", getMenu)

	e.Logger.Fatal(e.Start(":14045"))

}
