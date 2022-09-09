package main

import (
	"github/koybigino/getting-started-fiber/routers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routers.RouterPost(app)
	routers.RouterUser(app)

	log.Fatal(app.Listen(":3000"))
}
