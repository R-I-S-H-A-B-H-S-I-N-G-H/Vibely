package service

import (
	"errors"
	"fmt"

	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/dao"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/dto"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/entity"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/enum"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/mapper"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/utils"
)

type SongService struct{}

var songDao *dao.SongDAO
var songMapper *mapper.SongMapper
var s3Util *S3Service
var pathService *PathService
var audioProcessService *AudioProcessService

const S3_LINK_EXP = 60 * 60

func (s *SongService) getSongDao() *dao.SongDAO {
	if songDao == nil {
		songDao = dao.NewSongDAO()
	}
	return songDao
}

func (s *SongService) Get(id string) (*dto.SongDTO, error) {
	song, err := s.GetEntity(id)
	return songMapper.ToDTO(song), err
}

func (s *SongService) GetEntity(id string) (*entity.Song, error) {
	dao := s.getSongDao()
	return dao.FindByID(id)
}

func (s *SongService) ProcessSong(songShortId string) (string, error) {
	return audioProcessService.EncodeAudioToHLS(songShortId, 10, 28)
}

func (s *SongService) Save(dto *dto.SongDTO) (*dto.SongDTO, error) {
	dao := s.getSongDao()
	song := songMapper.FromDTO(dto, &entity.Song{})
	song, err := dao.Create(song)

	if err != nil {
		return nil, err
	}

	dto = songMapper.ToDTO(song)
	songPath := pathService.GetRawAudioS3Path(dto.ShortId)
	dto.PresignedUrl, err = s3Util.GeneratePresignedPutURL(songPath, S3_LINK_EXP)
	if err != nil {
		return nil, err
	}

	return dto, err
}

func (s *SongService) Update(id string, dto *dto.SongDTO) (*dto.SongDTO, error) {
	dao := s.getSongDao()
	song, err := s.GetEntity(id)

	if err != nil {
		return nil, err
	}

	song = songMapper.FromDTO(dto, song)
	err = dao.Update(song)
	return songMapper.ToDTO(song), err
}

func (s *SongService) GetList(filter map[string]interface{}, pageSize int, page int) ([]*dto.SongDTO, error) {
	dao := s.getSongDao()
	songList, err := dao.FindAll(filter, pageSize, page)
	songs := make([]*entity.Song, len(songList.Data))
	for i := range songList.Data {
		songs[i] = &songList.Data[i]
	}
	songsDTOList := songMapper.ToDTOList(songs)
	return songsDTOList, err
}

func (s *SongService) CanUpdateStatus(dto *dto.SongDTO, status enum.SongStatus) (bool, error) {
	if status == enum.StatusUploaded {
		songPath := pathService.GetRawAudioS3Path(dto.ShortId)
		return s3Util.ObjectExists(songPath)
	}
	return true, nil
}

func (s *SongService) UpdateStatus(id string, status enum.SongStatus) error {
	songDto, err := s.Get(id)
	if err != nil {
		return err
	}

	canupdate, err := s.CanUpdateStatus(songDto, status)
	if err != nil {
		return err
	}
	if !canupdate {
		return errors.New("status Cant be updated")
	}

	songDto.Status = status
	_, err = s.Update(id, songDto)
	return err
}

func (s *SongService) GenerateHLSJob() {
	// get all songs with status uploaded
	dao := s.getSongDao()
	res, err := dao.FindByStatus(enum.StatusUploaded, 10)
	if err != nil {
		return
	}

	songist := res.Data
	songs := make([]*entity.Song, len(songist))
	for i := range songist {
		songs[i] = &songist[i]
	}

	songDTOList := songMapper.ToDTOList(songs)
	for _, songDTO := range songDTOList {
		s.UpdateStatus(songDTO.ID, enum.StatusProcessing)
		songDTO := s.GenerateHLSForSong(songDTO)
		s.Update(songDTO.ID, songDTO)
	}
}

func (s *SongService) GenerateHLSForSong(songDTO *dto.SongDTO) *dto.SongDTO {
	stringObj, _ := utils.ToString(songDTO)
	fmt.Println()
	fmt.Println("GENERATE HLS FOR SONG :: ", stringObj)
	fmt.Println()

	type HLSVariant struct {
		SegmentDuration int
		BitrateKbps     int
		Bandwidth       int // in bps
	}

	variants := []HLSVariant{
		{2, 32, 32000},
		{2, 64, 64000},
		{3, 96, 96000},
		{5, 128, 128000},
		{5, 160, 160000},
		{10, 192, 192000},
		{10, 320, 320000},
	}

	for _, v := range variants {
		_, err := audioProcessService.EncodeAudioToHLS(songDTO.ShortId, v.SegmentDuration, v.BitrateKbps)

		if err != nil {
			fmt.Println("ERROR WHILE PROCESSING :: ", err.Error())
			continue
		}

		if songDTO.HLSStreams == nil {
			songDTO.HLSStreams = dto.ResolutionMapDTO{}
		}

		songDTO.HLSStreams[v.BitrateKbps] = dto.HLSStreamDTO{
			URL:       pathService.GetHLSAudioPlaylistS3Path(songDTO.ShortId, v.BitrateKbps),
			Bandwidth: uint(v.Bandwidth),
		}
	}

	songDTO.Status = enum.StatusProcessed
	return songDTO
}
