package middlewares

import (
	"fiber_first/util"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx)  error{
	cookie :=c.Cookies("jwt")

	if _,err := util.ParseJwt(cookie);err!=nil{
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message":"unauthenticated",
		})
	}
	//次の処理に映る
	return c.Next()
}