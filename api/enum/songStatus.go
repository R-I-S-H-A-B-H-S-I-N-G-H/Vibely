package enum

import "errors"

type SongStatus string

const (
	StatusPendingUpload SongStatus = "PENDING_UPLOAD" // just created, no upload yet
	StatusUploading     SongStatus = "UPLOADING"      // client started uploading
	StatusUploaded      SongStatus = "UPLOADED"       // upload complete
	StatusProcessing    SongStatus = "PROCESSING"     // backend is processing the audio
	StatusProcessed     SongStatus = "PROCESSED"      // ready to play
	StatusFailed        SongStatus = "FAILED"         // processing failed
)

// create a dunction to parese enum
func ParseSongStatus(status string) (SongStatus, error) {
	switch status {
	case "PENDING_UPLOAD":
		return StatusPendingUpload, nil
	case "UPLOADING":
		return StatusUploading, nil
	case "UPLOADED":
		return StatusUploaded, nil
	case "PROCESSING":
		return StatusProcessing, nil
	case "PROCESSED":
		return StatusProcessed, nil
	case "FAILED":
		return StatusFailed, nil
	default:
		return "", errors.New("invalid string")
	}
}
