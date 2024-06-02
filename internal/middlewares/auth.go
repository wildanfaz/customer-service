package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/go-market/configs"
	"github.com/wildanfaz/go-market/internal/helpers"
	"github.com/wildanfaz/go-market/internal/pkg"
	"github.com/wildanfaz/go-market/internal/repositories"
)

func Auth(log *logrus.Logger, config *configs.Config, usersRepo repositories.Users) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var resp = helpers.NewResponse()

		bearerToken := c.Get("Authorization")
		if bearerToken == "" || len(strings.Split(bearerToken, " ")) < 2 {
			log.Errorln("Invalid token")
			return c.Status(fiber.StatusUnauthorized).JSON(resp.AsError().WithMessage("Unauthorized"))
		}

		token := strings.Split(bearerToken, " ")[1]

		claims, err := pkg.ValidateToken(token, config.JWTSecretKey)
		if err != nil {
			log.Errorln(err.Error())
			return c.Status(fiber.StatusUnauthorized).JSON(resp.AsError().WithMessage("Unauthorized"))
		}

		user, err := usersRepo.GetUserByEmail(c.Context(), claims.Email)
		if err != nil {
			log.Errorln(err.Error())
			return c.Status(fiber.StatusUnauthorized).JSON(resp.AsError().WithMessage("Unauthorized"))
		}

		c.Locals("user", user)

		return c.Next()
	}
}