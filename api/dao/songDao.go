package dao

import (
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/dto"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/entity"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/enum"
)

type SongDAO struct {
	*DAO[entity.Song]
}

// NewSongDAO creates a new SongDAO
func NewSongDAO() *SongDAO {
	return &SongDAO{
		DAO: GetDAO[entity.Song](),
	}
}

// FindByStatus fetches all songs by a particular status
func (s *SongDAO) FindByStatus(status enum.SongStatus, size int) (*dto.PaginationDTO[entity.Song], error) {
	// Filters by status
	filters := map[string]interface{}{
		"status": status,
	}
	return s.FindAll(filters, 1, size)
}

// make function to find by shortId
func (s *SongDAO) FindByShortId(shortId string) (*entity.Song, error) {
	return s.FindOne(map[string]interface{}{
		"short_id": shortId,
	})
}
