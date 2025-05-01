package mapper

import (
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/dto"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/entity"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/utils"
)

type SongMapper struct {
}

func (m *SongMapper) ToDTO(song *entity.Song) *dto.SongDTO {
	songDTO := &dto.SongDTO{}
	songDTO.ID = song.ID
	songDTO.ShortId =song.ShortId
	songDTO.CreatedAt = song.CreatedAt
	songDTO.UpdatedAt = song.UpdatedAt
	songDTO.IsDeleted = song.IsDeleted

	songDTO.Title = song.Title
	songDTO.Description = song.Description
	songDTO.Artist = song.Artist
	songDTO.Album = song.Album
	songDTO.Genres = song.Genres
	songDTO.Language = song.Language
	songDTO.Duration = song.Duration
	songDTO.ReleaseDate = song.ReleaseDate

	var hlsStreams dto.ResolutionMapDTO

	err := utils.ToObject(song.HLSStreams, &hlsStreams)
	if err != nil {
		panic(err)
	}
	songDTO.HLSStreams = hlsStreams

	return songDTO
}

func (m *SongMapper) FromDTO(songDTO *dto.SongDTO, song *entity.Song) *entity.Song {
	if song == nil {
		song = &entity.Song{}
	}

	song.Title = songDTO.Title
	song.Description = songDTO.Description
	song.Artist = songDTO.Artist
	song.Album = songDTO.Album
	song.Genres = songDTO.Genres
	song.Language = songDTO.Language
	song.Duration = songDTO.Duration
	song.ReleaseDate = songDTO.ReleaseDate

	var err error
	song.HLSStreams, err = utils.ToString(songDTO.HLSStreams)

	if err != nil {
		panic(err)
	}

	return song
}
