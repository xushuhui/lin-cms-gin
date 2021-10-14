package biz

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"lin-cms-go/internal/data/ent"
	"lin-cms-go/pkg/utils"
)

func LocalUser(c *fiber.Ctx) (user ent.LinUser) {
	jwtToken := c.Locals("user").(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	bytes, _ := utils.JSONEncode(claims["user"])
	utils.JSONDecode(bytes, &user)
	return
}
