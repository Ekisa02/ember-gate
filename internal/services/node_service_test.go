package services

import (
	"testing"

	"github.com/example/ember-gate/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	if err := db.AutoMigrate(&models.Node{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}

func TestNodeCreateAndList(t *testing.T) {
	db := setupDB(t)
	svc := NewNodeService(db)

	node := &models.Node{Name: "Test Node", URL: "https://node.example.org", Provider: "local"}
	if err := svc.Create(node); err != nil {
		t.Fatalf("create failed: %v", err)
	}

	list, err := svc.List()
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("expected 1 node, got %d", len(list))
	}
}
