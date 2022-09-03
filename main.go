package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	Id        int    `json:"id" gorm:"primaryKey"`
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

	dsn := "host=localhost user=postgres password=Bielem@*01 dbname=fiber_bd port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error to connect to our databse : %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("connection to the database ok .... %v", db)

	app.Get("/posts", func(c *fiber.Ctx) error {
		var post []Post
		result := db.Find(&post)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Error to get all elements of the table",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "All post are getting well",
			"data":    post,
		})
	})

	app.Get("/posts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		post := new(Post)
		intId, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Please enter a integer like parameter !",
			})
		}
		e := db.First(post, intId).Error
		if e != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "post not found !",
				"error":   e,
			})
		}
		return c.JSON(post)
	})

	app.Post("/posts", func(c *fiber.Ctx) error {
		newPost := new(Post)

		if err := c.BodyParser(newPost); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		err := db.Create(newPost).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Your post wasn't been created !",
				"error":   err,
			})
		}
		return c.JSON(fiber.Map{
			"message": "Your post was been created !",
			"data":    newPost,
		})
	})

	app.Delete("/posts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		post := new(Post)
		intId, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Please enter a integer like parameter !",
			})
		}
		db.Delete(post, intId)
		return c.SendStatus(fiber.StatusNoContent)
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

		e := db.First(newPost, intId)
		if e.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "post not found !",
			})
		}

		newPost.Title = "New Title update"
		newPost.Content = "New Content update"
		db.Save(newPost)

		return c.JSON(newPost)

	})

	log.Fatal(app.Listen(":3000"))
}
