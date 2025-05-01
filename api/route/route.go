package route

import "github.com/gofiber/fiber/v2"

type Route struct{}

var userRoute *UserRouter

func (r *Route) HandleRoute(app *fiber.App) {
	apiV1 := app.Group("/api")

	userRoute.handleRoute(apiV1)
}
