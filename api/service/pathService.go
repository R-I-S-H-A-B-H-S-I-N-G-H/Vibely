package service

import (
	"fmt"
	"net/url"
	"os"
)

type PathService struct{}

func NewPathService() *PathService {
	return &PathService{}
}

func (p *PathService) GetSelfServicePath() string {
	return os.Getenv("SELF")
}

func (p *PathService) GetSongS3BaseFolder() string {
	return "songs"
}

func (p *PathService) GetS3BaseUrl() string {
	return os.Getenv("S3_BASE_URL")
}

func (p *PathService) GetRawAudioS3Path(songShortId string) string {
	return fmt.Sprintf("%s/%s/raw", p.GetSongS3BaseFolder(), songShortId)
}

func (p *PathService) GetFullRawAudioS3Path(songShortId string) string {
	return fmt.Sprintf("%s/%s", p.GetS3BaseUrl(), p.GetRawAudioS3Path(songShortId))
}

func (p *PathService) GetHLSAudioS3Path(songShortId string, bitrate int) string {
	return fmt.Sprintf("%s/%s/hls/%d", p.GetSongS3BaseFolder(), songShortId, bitrate)
}

func (p *PathService) GetFullHLSAudioS3Path(songShortId string, bitrate int) string {
	return fmt.Sprintf("%s/%s", p.GetS3BaseUrl(), p.GetHLSAudioS3Path(songShortId, bitrate))
}

func (p *PathService) GetHLSAudioPlaylistS3Path(songShortId string, bitrate int) string {
	return fmt.Sprintf("%s/%s/hls/%d/playlist.m3u8", p.GetSongS3BaseFolder(), songShortId, bitrate)
}

func (p *PathService) GetHLSAudioMasterPlaylistS3Path(songShortId string) string {
	return fmt.Sprintf("%s/%s/hls/playlist.m3u8", p.GetSongS3BaseFolder(), songShortId)
}

func (p *PathService) GetFullHLSAudioPlaylistS3Path(songShortId string, bitrate int) string {
	return fmt.Sprintf("%s/%s", p.GetS3BaseUrl(), p.GetHLSAudioPlaylistS3Path(songShortId, bitrate))
}

func (p *PathService) GetLambdaCallback(service string, id string, queryParams map[string]string) string {
	basePath := fmt.Sprintf("%s/api/%s/%s/cb", p.GetSelfServicePath(), service, id)

	if len(queryParams) == 0 {
		return basePath
	}

	q := url.Values{}
	for key, value := range queryParams {
		q.Set(key, value)
	}

	return fmt.Sprintf("%s?%s", basePath, q.Encode())
}
