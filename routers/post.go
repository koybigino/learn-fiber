package routers

import (
	"github/koybigino/getting-started-fiber/middleware"
	"github/koybigino/getting-started-fiber/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RouterPost(router *fiber.App) {
	p := router.Group("/posts/", middleware.AuthRequired())
	p.Get("", GetAllPost)
	p.Get(":id", GetPost)
	p.Post("", CreatePost)
	p.Delete(":id", DeletePost)
	p.Put(":id", UpdatePost)
}

func GetAllPost(c *fiber.Ctx) error {
	var post []models.Post
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
}

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	post := new(models.Post)
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
}

func CreatePost(c *fiber.Ctx) error {
	newPost := new(models.Post)

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
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	post := new(models.Post)
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Please enter a integer like parameter !",
		})
	}
	db.Delete(post, intId)
	return c.SendStatus(fiber.StatusNoContent)
}

func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	newPost := new(models.Post)
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
}
