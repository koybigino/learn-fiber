package main

import (
	"github/koybigino/getting-started-fiber/routers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "*",
		AllowMethods:     "*",
		AllowCredentials: true,
	}))

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World !",
		})
	})

	routers.RouterAuth(app)

	routers.RouterPost(app)
	routers.RouterUser(app)
	routers.RouterVote(app)

	log.Fatal(app.Listen(":3000"))
}
