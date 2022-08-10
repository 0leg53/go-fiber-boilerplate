package database

import (
	"fmt"

	"app/configs"
	"app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Its a global variable!
var DB *gorm.DB

func Connect() {
	var (
		host     = "db"
		port     = 5432
		user     = configs.Env("POSTGRES_USER")
		password = configs.Env("POSTGRES_PASSWORD")
		dbname   = configs.Env("POSTGRES_DB")
	)
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("❌❌❌ Couldn't connect to database")
	}
	fmt.Println("Successfully connected to psql ✔️ 🐘")

	DB = connection

	connection.AutoMigrate(&models.User{})
}
