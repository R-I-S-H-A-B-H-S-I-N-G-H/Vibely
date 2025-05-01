package dto

import "time"

type BaseDTO struct {
	ID        string    `json:"id"`
	ShortId   string    `json:"shortId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsDeleted bool      `json:"isDeleted"`
}
