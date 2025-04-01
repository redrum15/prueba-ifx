package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VirtualMachine struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"not null"`
	Cores     int       `gorm:"not null"`
	Ram       int       `gorm:"not null"`
	OS        string    `gorm:"not null"`
	Status    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
