package dto

type HLSStreamDTO struct {
	URL       string `json:"url"`
	Bandwidth uint   `json:"bandwidth"` // in kbps
	Codec     string `json:"codec"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

type ResolutionMapDTO map[int]HLSStreamDTO
