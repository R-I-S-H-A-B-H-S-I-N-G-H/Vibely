package route

import (
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/controller"
	"github.com/gofiber/fiber/v2"
)

type SongRoute struct{}

var songController *controller.SongController

func (u *SongRoute) handleRoute(app fiber.Router) {
	userRoute := app.Group("/song")

	userRoute.Post("", songController.CreateSong)
	userRoute.Put("", songController.UpdateSong)
}
