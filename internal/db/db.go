package db

import (
	"fmt"

	"github.com/example/ember-gate/internal/config"
	"github.com/example/ember-gate/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg *config.Config) (*gorm.DB, error) {
	if cfg.DBDSN == "" {
		return nil, fmt.Errorf("DB_DSN not provided")
	}
	dialector := postgres.Open(cfg.DBDSN)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate for development
	if err := db.AutoMigrate(&models.User{}, &models.Node{}); err != nil {
		return nil, err
	}

	return db, nil
}
