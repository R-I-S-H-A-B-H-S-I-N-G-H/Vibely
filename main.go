package main

import (
	databaseconfig "github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/database-config"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/route"
	"github.com/gofiber/fiber/v2"
)

var mainRouter *route.Route
var database *databaseconfig.Database

func main() {

	// initalizing db
	_, err := database.Init()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	mainRouter.HandleRoute(app)

	app.Listen(":3000")

}
