package database

import (
	"gorm.io/plugin/opentelemetry/tracing"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB(dbAddress string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	if err := db.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
	seedDB(db)
	return db
}
