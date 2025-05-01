package route

import (
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/controller"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct{}

var userController *controller.UserController

func (u *UserRouter) handleRoute(app fiber.Router) {
	userRoute := app.Group("/user")

	userRoute.Get("/:id", userController.GetUser)
}
