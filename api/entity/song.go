package entity

import (
	"time"
)

type Song struct {
	BaseEntity
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Artist      string `gorm:"type:varchar(255);not null"`
	Album       string `gorm:"type:varchar(255)"`
	Genres      string `gorm:"type:varchar(255)"`
	Language    string `gorm:"type:char(2);default:'en'"`
	Duration    uint   `gorm:"not null"`
	ReleaseDate time.Time

	HLSStreams string `gorm:"type:longtext"`
}
