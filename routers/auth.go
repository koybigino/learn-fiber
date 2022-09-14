package routers

import (
	"github/koybigino/getting-started-fiber/models"
	"github/koybigino/getting-started-fiber/oauth"
	"github/koybigino/getting-started-fiber/session"
	"github/koybigino/getting-started-fiber/utils"

	"github.com/gofiber/fiber/v2"
)

func RouterAuth(router *fiber.App) {
	router.Post("/login", Login)
}

func Login(c *fiber.Ctx) error {

	user := new(models.User)
	body := new(models.User)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}
	e := db.Where("email = ?", string(body.Email)).First(user).Error
	if e != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": e.Error(),
		})
	}

	//Store the current user in the session
	sess := session.CreateNewSession(c)
	sess.Set("current_user_id", user.Id)
	sess.Set("current_user_email", user.Email)
	if err := sess.Save(); err != nil {
		panic(err.Error())
	}

	// if body.Password != user.Password {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "Invalid credentials",
	// 	})
	// }

	if utils.Verify(body.Password, user.Password) != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	token := oauth.CreateAccessToken()

	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	return c.JSON(fiber.Map{
		"token":      token,
		"token_type": "Bearer",
	})
}
