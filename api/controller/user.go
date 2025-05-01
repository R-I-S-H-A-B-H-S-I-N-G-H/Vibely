package controller

import (
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/dao"
	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/entity"
	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

var userDao *dao.DAO[entity.User]

func (u *UserController) getUserDao() *dao.DAO[entity.User] {
	if userDao == nil {
		userDao = dao.GetDAO[entity.User]()
	}
	return userDao
}

func (u *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	u.getUserDao().Create(&entity.User{
		Name: "fifi" + id,
	})
	return c.SendString("User ID: " + id)
}
