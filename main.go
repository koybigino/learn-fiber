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

	app.Get("/posts", func(c *fiber.Ctx) error {
		return c.JSON(postList)
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

	app.Delete("/posts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Please enter a integer like parameter !",
			})
		}
		for index, post := range postList {
			if post.Id == intId {
				postList = append(postList[:index], postList[index+1:]...)
				return c.SendStatus(fiber.StatusNoContent)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Post not found to delete it",
		})
	})

	app.Put("/posts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		intId, err := strconv.Atoi(id)
		newPost := new(Post)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Please enter a integer like parameter",
			})
		}

		for index, post := range postList {
			if intId == post.Id {
				if err := c.BodyParser(newPost); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"message": err.Error(),
					})
				}
				postList = append(postList[:index], postList[index+1:]...)
				newPost.Id = intId
				postList = append(postList, *newPost)
				c.SendStatus(fiber.StatusCreated)
				return c.JSON(postList)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Post not found to update it !",
		})

	})

	log.Fatal(app.Listen(":3000"))
}
