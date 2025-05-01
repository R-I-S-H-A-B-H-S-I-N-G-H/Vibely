package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseEntity defines common fields and behavior
type BaseEntity struct {
	ID        string    `gorm:"primaryKey;type:char(36)"` // Use CHAR(36) for MySQL UUID
	ShortId   string    `gorm:"unique;type:char(10)"`     // Use CHAR(36) for MySQL UUID
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	IsDeleted bool      `gorm:"default:false"`
}

// BeforeCreate hook to generate a UUID for the ID field
func (b *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" { // Only generate if not already set
		b.ID = uuid.New().String()
	}

	if b.ShortId == "" {
		b.ShortId = uuid.New().String()[0:8]
	}
	return
}
