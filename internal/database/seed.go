package database

import (
	"fmt"
	"golang-api-restaurant/internal/model"
	"golang-api-restaurant/internal/model/constant"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	// migrate the schema
	db.AutoMigrate(&model.MenuItem{}, &model.Order{}, &model.ProductOrder{})
	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Rica Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      constant.MenuTypeFood,
		},
	}

	drinksMenu := []model.MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     4000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Es Teh Manis",
			OrderCode: "es_teh_manis",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		},
	}

	fmt.Println(foodMenu)
	fmt.Println(drinksMenu)

	var count int64
	db.Model(&model.MenuItem{}).Count(&count)
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

}
