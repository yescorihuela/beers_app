package config

import (
	"testing"

	"gorm.io/gorm"
)

func TestGetDatabaseURL(t *testing.T) {
	expected := "postgres://user_challenge:pa55word@localhost:5432/code_challenge?sslmode=disable"
	got := getDatabaseURL()
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
