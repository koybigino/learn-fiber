package routers

import (
	"github/koybigino/getting-started-fiber/databases"
	"github/koybigino/getting-started-fiber/middleware"
	"github/koybigino/getting-started-fiber/models"
	"github/koybigino/getting-started-fiber/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var db = databases.Connection()

func RouterUser(router *fiber.App) {
	u := router.Group("/users/")
	u.Get(":id", GetUser, middleware.AuthRequired())
	u.Post("", CreateUser)
}

func CreateUser(c *fiber.Ctx) error {
	newUser := new(models.User)

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(*newUser)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	if pwd, e := utils.Hash(newUser.Password); e != nil {
		return c.SendString("Error to hash password")
	} else {
		newUser.Password = string(pwd)
	}

	err := db.Create(newUser).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Your user wasn't been created !",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Your user was been created !",
		"data":    newUser,
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	db.Preload("Posts").Preload("Votes").Find(user, intId)
	// e := db.First(user, intId).Error
	// if e != nil {
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"error": e.Error(),
	// 	})
	// }
	return c.JSON(user)
}
