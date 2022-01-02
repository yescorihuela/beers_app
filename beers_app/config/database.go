package config

import (
	"fmt"
	"log"
	"os"

	"github.com/yescorihuela/beers_app/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDatabaseURL() string {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
}

func ConnectDatabase() *gorm.DB {
	dbURL := getDatabaseURL()
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&domain.Beer{})
	return db
}
