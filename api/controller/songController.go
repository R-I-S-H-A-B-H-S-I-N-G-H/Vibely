package controller

import (
	"strconv"

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

func (s *SongController) GetList(c *fiber.Ctx) error {
	filters := make(map[string]interface{}) // Add appropriate filters if needed
	// page := 1                               // Set the desired page number
	// pageSize := 10                          // Set the desired page size

	page, err := strconv.ParseInt(c.Query("page", "1"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid page number",
		})
	}

	pageSize, err := strconv.ParseInt(c.Query("size", "10"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid page number",
		})
	}

	songDTOs, err := songService.GetList(filters, int(page), int(pageSize))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	c.JSON(songDTOs)
	return err
}

func (s *SongController) ProcessSong(c *fiber.Ctx) error {
	songShortId := c.Params("id")
	res, err := songService.ProcessSong(songShortId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}
	// c.JSON(songDTO)
	return c.Status(fiber.StatusOK).JSON(res)
}
