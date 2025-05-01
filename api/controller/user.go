package controller

import "github.com/gofiber/fiber/v2"

type UserController struct{}

func (u *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("User ID: " + id)
}
