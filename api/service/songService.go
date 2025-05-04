package service

import (
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/dao"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/dto"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/entity"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/mapper"
)

type SongService struct{}

var songDao *dao.DAO[entity.Song]
var songMapper *mapper.SongMapper
var s3Util *S3Service
var pathService *PathService
var audioProcessService *AudioProcessService

const S3_LINK_EXP = 60 * 60

func (s *SongService) getSongDao() *dao.DAO[entity.Song] {
	if songDao == nil {
		songDao = dao.GetDAO[entity.Song]()
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
