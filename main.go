package main

import (
	databaseconfig "github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/database-config"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/route"
	"github.com/gofiber/fiber/v2"
)

var mainRouter *route.Route
var database *databaseconfig.Database

func main() {
	const PORT = ":4001"
	// initalizing db
	_, err := database.Init()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	mainRouter.HandleRoute(app)

	println("Starting server at port::  " + PORT)
	err = app.Listen(PORT)
	if err != nil {
		panic(err)
	}
}
