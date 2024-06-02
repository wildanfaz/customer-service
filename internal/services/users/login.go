package users

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wildanfaz/go-market/internal/helpers"
	"github.com/wildanfaz/go-market/internal/models"
	"github.com/wildanfaz/go-market/internal/pkg"
)

func (s *ImplementServices) Login(c *fiber.Ctx) error {
	var (
		resp    = helpers.NewResponse()
		payload models.User
	)

	err := c.BodyParser(&payload)
	if err != nil {
		s.log.Errorln("Login :", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().WithMessage("Invalid data"))
	}

	err = validation.ValidateStruct(&payload,
		validation.Field(&payload.Email, validation.Required, is.Email),
		validation.Field(&payload.Password, validation.Required),
	)
	if err != nil {
		s.log.Errorln("Login :", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().WithMessage(err.Error()))
	}

	user, err := s.usersRepo.GetUserByEmail(c.Context(), payload.Email)
	if err != nil {
		s.log.Errorln("Login :", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().WithMessage("Invalid email or password"))
	}

	err = pkg.ComparePassword(payload.Password, user.Password)
	if err != nil {
		s.log.Errorln("Login :", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().WithMessage("Invalid email or password"))
	}

	claims := pkg.NewClaims{Email: user.Email}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(s.config.JWTDuration))

	token, err := pkg.GenerateToken(&claims, s.config.JWTSecretKey)
	if err != nil {
		s.log.Errorln("Login :", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(resp.AsError().WithMessage("Internal server error"))
	}

	s.log.Infoln("Login success")
	return c.Status(fiber.StatusOK).JSON(resp.WithMessage("Login success").
	WithData(map[string]string{"token": token}))
}