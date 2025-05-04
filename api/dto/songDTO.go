package dto

import (
	"time"

	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/enum"
)

type SongDTO struct {
	BaseDTO
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Artist      string    `json:"artist"`
	Album       string    `json:"album"`
	Genres      string    `json:"genres"`
	Language    string    `json:"language"`
	Duration    uint      `json:"duration"`
	ReleaseDate time.Time `json:"releaseDate"`

	HLSStreams   ResolutionMapDTO `json:"hlsStreams"`
	PresignedUrl string           `json:"presignedUrl"`
	Status       enum.SongStatus  `json:"status"`
}
