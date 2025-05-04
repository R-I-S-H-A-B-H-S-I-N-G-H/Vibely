package entity

import (
	"time"

	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/enum"
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

	HLSStreams string          `gorm:"type:longtext"`
	Status     enum.SongStatus `gorm:"type:varchar(20);not null;default:'PENDING_UPLOAD'"`
}
