package routers

import (
	"github/koybigino/getting-started-fiber/middleware"
	"github/koybigino/getting-started-fiber/models"
	"github/koybigino/getting-started-fiber/session"
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
	result := db.Find(&post).Limit(10)
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
	// e := db.First(post, intId).Error
	// if e != nil {
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"message": "post not found !",
	// 		"error":   e,
	// 	})
	// }
	db.Preload("Votes").First(post, intId)
	// db.Joins("Company").Joins("Manager").Joins("Account").First(&user, "users.name = ?", "jinzhu")
	// db.Joins("Company").Joins("Manager").Joins("Account").Find(&users, "users.id IN ?", []int{1,2,3,4,5})

	return c.JSON(post)
}

func CreatePost(c *fiber.Ctx) error {
	newPost := new(models.Post)
	sess := session.CreateNewSession(c)
	currentUserId := sess.Get("current_user_id")

	if err := c.BodyParser(newPost); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	newPost.UserId = currentUserId.(int)
	err := db.Create(newPost).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Your post wasn't been created !",
			"error":   err.Error(),
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
	sess := session.CreateNewSession(c)
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Please enter a integer like parameter !",
		})
	}
	currentUserId := sess.Get("current_user_id")
	// currentUserEmail := sess.Get("current_user_email")

	e := db.First(post, intId).Error
	if e != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "post not found !",
			"error":   e.Error(),
		})
	}

	if currentUserId != post.UserId {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
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
