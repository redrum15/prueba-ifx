package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Type      string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (vm *VirtualMachine) BeforeCreate(tx *gorm.DB) (err error) {
	vm.ID = uuid.New()
	return nil
}
