package models

import (
	"time"

	"gorm.io/gorm"
)

type Node struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Name       string         `gorm:"size:255;not null" json:"name"`
	URL        string         `gorm:"size:1024;not null" json:"url"`
	Provider   string         `gorm:"size:255" json:"provider"`
	Health     string         `gorm:"size:64;default:'unknown'" json:"health"`
	OwnerID    uint           `json:"owner_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
