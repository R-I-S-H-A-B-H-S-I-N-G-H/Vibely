package controller

import (
	"strconv"

	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/dto"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/enum"
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

func (s *SongController) UpdateSongStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	status := c.Params("status")
	statusEnum, err := enum.ParseSongStatus(status)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid status value",
		})
	}

	err = songService.UpdateStatus(id, statusEnum)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Status updated successfully",
	})
}

func (s *SongController) LambdaCallback(c *fiber.Ctx) error {
	shortId := c.Params("shortId")
	segmentStr := c.Query("segment")
	bitrateStr := c.Query("bitrate")
	bandwidthStr := c.Query("bandwidth")

	segment, err := strconv.Atoi(segmentStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid segment value",
		})
	}

	bitrate, err := strconv.Atoi(bitrateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid bitrate value",
		})
	}

	bandwidth, err := strconv.Atoi(bandwidthStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid bandwidth value",
		})
	}

	var lambdaCb *dto.LambdaCallbackResponse
	c.BodyParser(&lambdaCb)

	songService.LambdaCallbackHandler(shortId, segment, bitrate, bandwidth, lambdaCb)
	return nil
}
