package dto

type LambdaPayloadDTO struct {
	ResourcePath string `json:"resource_path"`
	S3UpPath     string `json:"s3_up_path"`
	SegmentTime  int    `json:"segment_time"`
	AudioBitrate string `json:"audio_bitrate"`
}
