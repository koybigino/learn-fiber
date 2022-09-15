package session

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	// "github.com/gofiber/storage/sqlite3"
)

// var storage = sqlite3.New()
// var store = session.New(session.Config{
// 	Storage: storage,
// })
var store = session.New()

func CreateNewSession(c *fiber.Ctx) *session.Session {
	sess, err := store.Get(c)
	if err != nil {
		panic(err.Error())
	}

	return sess
}
