package oauth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateAccessToken() string {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	//Create the Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ""
	}

	return t
}

// func VerifyAccessToken(c *fiber.Ctx) string {
// 	user := c.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	email := claims["sub"].(string)

// 	return email
// }
