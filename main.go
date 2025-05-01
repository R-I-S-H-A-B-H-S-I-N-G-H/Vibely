package main

import (
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/route"
	"github.com/gofiber/fiber/v2"
)

var mainRouter *route.Route

func main() {
	app := fiber.New()

	mainRouter.HandleRoute(app)

	app.Listen(":3000")

}
