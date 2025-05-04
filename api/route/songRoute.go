package route

import (
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/controller"
	"github.com/gofiber/fiber/v2"
)

type SongRoute struct{}

var songController *controller.SongController

func (u *SongRoute) handleRoute(app fiber.Router) {
	songRoute := app.Group("/song")

	songRoute.Post("", songController.CreateSong)
	songRoute.Put("", songController.UpdateSong)
	songRoute.Get("/list", songController.GetList)
	songRoute.Get("/:id", songController.ProcessSong)
	songRoute.Get("/:id/:status", songController.UpdateSongStatus)
}
