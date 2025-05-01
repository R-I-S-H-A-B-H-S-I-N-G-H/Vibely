package route

import "github.com/gofiber/fiber/v2"

type Route struct{}

var userRoute *UserRouter
var songRoute *SongRoute

func (r *Route) HandleRoute(app *fiber.App) {
	apiV1 := app.Group("/api")

	userRoute.handleRoute(apiV1)
	songRoute.handleRoute(apiV1)
}
