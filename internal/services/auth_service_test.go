package services

import (
	"testing"

	"github.com/example/ember-gate/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}

func TestAuthRegisterAndAuthenticate(t *testing.T) {
	db := setupTestDB(t)
	svc := NewAuthService(db, "secret")

	user, err := svc.Register("bob@example.com", "password123")
	if err != nil {
		t.Fatalf("register failed: %v", err)
	}
	if user.Email != "bob@example.com" {
		t.Fatalf("unexpected email: %s", user.Email)
	}

	// Authenticate
	token, err := svc.Authenticate("bob@example.com", "password123")
	if err != nil {
		t.Fatalf("authenticate failed: %v", err)
	}
	if token == "" {
		t.Fatalf("empty token")
	}
}
