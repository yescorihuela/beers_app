package config

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/gorm"
)

func TestGetDatabaseURL(t *testing.T) {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	expected := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	got := GetDatabaseURL()
	if got != expected {
		t.Errorf("Expected: %s, got: %s", expected, got)
	}
}

func TestConnectDatabase(t *testing.T) {
	var expected *gorm.DB
	got := ConnectDatabase()
	if got == expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}
