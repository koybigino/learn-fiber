package routers

import (
	"github/koybigino/getting-started-fiber/middleware"
	"github/koybigino/getting-started-fiber/models"
	"github/koybigino/getting-started-fiber/session"

	"github.com/gofiber/fiber/v2"
)

func RouterVote(router *fiber.App) {
	v := router.Group("/votes/", middleware.AuthRequired())
	v.Post("", CreateVote)
}

func CreateVote(c *fiber.Ctx) error {
	body := new(models.BodyVote)
	vote := new(models.Vote)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	sess := session.CreateNewSession(c)
	currentUserId := sess.Get("current_user_id")

	if e := db.Where("post_id = ? and user_id = ?", body.PostId, currentUserId.(int)).First(vote).Error; e != nil {
		vote.PostId = body.PostId
		vote.UserId = currentUserId.(int)
		err := db.Create(vote).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Your vote wasn't been created !",
				"error":   err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "Your vote was been created !",
			"data":    vote,
		})
	} else {
		db.Delete(vote, vote.PostId, vote.UserId)
		return c.SendStatus(fiber.StatusNoContent)
	}

}
