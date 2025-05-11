package dto

type LambdaCallbackPayload struct {
	Url    string            `json:"url"`
	Header map[string]string `json:"header"`
	Method string            `json:"method"`
}

type LambdaPayloadDTO struct {
	ResourcePath  string                  `json:"resource_path"`
	S3UpPath      string                  `json:"s3_up_path"`
	SegmentTime   int                     `json:"segment_time"`
	AudioBitrate  string                  `json:"audio_bitrate"`
	CallBackProps []LambdaCallbackPayload `json:"callback_props"`
}
