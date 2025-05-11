package dto

type HLSVariantDTO struct {
	SegmentDuration int
	BitrateKbps     int
	Bandwidth       int // in bps
	Codecs          string
}
