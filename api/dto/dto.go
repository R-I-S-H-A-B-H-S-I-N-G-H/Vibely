package dto

import "time"

type BaseDTO struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool
}
