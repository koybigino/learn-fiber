package main

import (
	"github/koybigino/getting-started-fiber/routers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	routers.RouterAuth(app)

	routers.RouterPost(app)
	routers.RouterUser(app)

	log.Fatal(app.Listen(":3000"))
}
