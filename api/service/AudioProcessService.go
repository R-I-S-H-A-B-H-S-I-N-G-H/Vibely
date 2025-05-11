package service

import (
	"fmt"

	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/dto"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/utils"
)

type AudioProcessService struct{}

var lambdaService *LambdaService

func NewAudioProcessService() *AudioProcessService {
	return &AudioProcessService{}
}

func (a *AudioProcessService) GetAudioBitrate(bitrateInKbps int) string {
	return fmt.Sprintf("%dk", bitrateInKbps)
}

func (a *AudioProcessService) EncodeAudioToHLS(songShortId string, segmentTime int, audioBitrateInKbps int, bandwidth int) (string, error) {
	queryParams := map[string]string{
		"segment":   fmt.Sprintf("%d", segmentTime),
		"bitrate":   fmt.Sprintf("%d", audioBitrateInKbps),
		"bandwidth": fmt.Sprintf("%d", bandwidth),
	}

	payload := dto.LambdaPayloadDTO{
		ResourcePath: pathService.GetFullRawAudioS3Path(songShortId),
		S3UpPath:     pathService.GetHLSAudioS3Path(songShortId, audioBitrateInKbps),
		SegmentTime:  segmentTime,
		AudioBitrate: a.GetAudioBitrate(audioBitrateInKbps),
		CallBackProps: []dto.LambdaCallbackPayload{
			{
				Url:    pathService.GetLambdaCallback("song", songShortId, queryParams),
				Header: map[string]string{},
				Method: "POST",
			},
		},
	}

	payloadString, err := utils.ToString(payload)
	if err != nil {
		panic(err)
	}

	lambdaPayload := map[string]string{
		"body": payloadString,
	}

	fmt.Println()
	fmt.Println("payload :: ", payloadString)
	fmt.Println()

	resp, err := lambdaService.InvokeLambda("Vibely-lam-py-HelloWorldFunction-xG7PQTAB8ldy", lambdaPayload)
	return string(resp), err
}
