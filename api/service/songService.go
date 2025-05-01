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

func (s *SongService) Save(dto *dto.SongDTO) (*dto.SongDTO, error) {
	dao := s.getSongDao()
	song := songMapper.FromDTO(dto, &entity.Song{})
	song, err := dao.Create(song)
	return songMapper.ToDTO(song), err
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
