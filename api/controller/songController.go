package controller

import (
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/dto"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/service"
	"github.com/gofiber/fiber/v2"
)

type SongController struct{}

var songService *service.SongService

func (s *SongController) GetSong(c *fiber.Ctx) error {
	return nil
}

func (s *SongController) CreateSong(c *fiber.Ctx) error {
	var songDTO *dto.SongDTO
	c.BodyParser(&songDTO)

	songDTO, err := songService.Save(songDTO)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	c.JSON(songDTO)
	return err
}

func (s *SongController) UpdateSong(c *fiber.Ctx) error {
	var songDTO *dto.SongDTO
	c.BodyParser(songDTO)

	songId := c.Params("id")

	songDTO, err := songService.Update(songId, songDTO)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	c.JSON(songDTO)
	return err
}
