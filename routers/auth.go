package routers

import (
	"github/koybigino/getting-started-fiber/models"
	"github/koybigino/getting-started-fiber/oauth"
	"github/koybigino/getting-started-fiber/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store = session.New()

func RouterAuth(router *fiber.App) {
	router.Post("/login", Login)
}

func Login(c *fiber.Ctx) error {

	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}

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
	sess.Set("current_user", user)

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
