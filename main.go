package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Post struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
}

var (
	postList = []Post{
		{Id: 1, Title: "title of post 1", Content: "small content of post 1 and some description", Published: true},
		{Id: 2, Title: "title of post 2", Content: "small content of post 2 and some description", Published: true},
	}
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/posts/:id", func(c *fiber.Ctx) error {
		getPost := new(Post)
		for _, post := range postList {
			if strconv.Itoa(post.Id) == c.Params("id") {
				getPost = &post
				break
			}
		}
		if getPost.Title == "" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Post not found !",
			})
		}
		return c.JSON(getPost)
	})

	app.Post("/posts", func(c *fiber.Ctx) error {
		newPost := new(Post)

		if err := c.BodyParser(newPost); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})

		}
		newPost.Id = rand.Intn(1000)
		postList = append(postList, *newPost)
		return c.JSON(postList)
	})
	log.Fatal(app.Listen(":3000"))
}
