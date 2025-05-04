package main

import (
	"os"

	databaseconfig "github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/database-config"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

var mainRouter *route.Route
var database *databaseconfig.Database

func main() {
	const PORT = ":4001"

	setupEnvVars()
	// initalizing db
	_, err := database.Init()
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	mainRouter.HandleRoute(app)

	println("Starting server at port::  " + PORT)
	err = app.Listen(PORT)
	if err != nil {
		panic(err)
	}
}

func setupEnvVars() {
	// Load .env file only in local development
	if os.Getenv("ENV") == "production" {
		return
	}
	err := godotenv.Load()
	if err != nil {
		println(err.Error())
		panic(err)
	}
}
