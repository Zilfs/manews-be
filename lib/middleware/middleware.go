package middleware

import (
	"manews/config"
	"manews/internal/adapter/handler/response"
	"manews/lib/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Middleware interface {
	CheckToken() fiber.Handler
}

type Options struct {
	authJwt auth.Jwt
}

// CheckToken implements Middleware.
func (o *Options) CheckToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHandler := c.Get("Authorization")
		var errorRespopnonse response.ErrorResponseDefault
		if authHandler == "" {
			errorRespopnonse.Meta.Status = false
			errorRespopnonse.Meta.Message = "Missing Authorization Header"
			return c.Status(fiber.StatusUnauthorized).JSON(errorRespopnonse)
		}

		tokenString := strings.Split(authHandler, "Bearer ")[1]
		claims, err := o.authJwt.VerifyAccessToken(tokenString)
		if err != nil {
			errorRespopnonse.Meta.Status = false
			errorRespopnonse.Meta.Message = "Invalid Token"
			return c.Status(fiber.StatusUnauthorized).JSON(errorRespopnonse)
		}
		c.Locals("user", claims)

		return c.Next()
	}
}

func NewMiddleware(cfg *config.Config) Middleware {
	opt := new(Options)
	opt.authJwt = auth.NewJwt(cfg)
	return opt
}
