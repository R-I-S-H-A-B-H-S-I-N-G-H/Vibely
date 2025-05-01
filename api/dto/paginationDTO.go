package dto

type PaginationDTO[T any] struct {
	Data       []T   `json:"data"`       // Paginated data
	TotalCount int64 `json:"totalCount"` // Total number of matching records
	Page       int64 `json:"page"`       // Current page
	Size       int64 `json:"size"`       // Page size
}
