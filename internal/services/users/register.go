package users

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofiber/fiber/v2"
	"github.com/wildanfaz/go-market/internal/helpers"
	"github.com/wildanfaz/go-market/internal/models"
	"github.com/wildanfaz/go-market/internal/pkg"
)

func (s *ImplementServices) Register(c *fiber.Ctx) error {
	var (
		resp = helpers.NewResponse()
		payload models.User
	)

	err := c.BodyParser(&payload)
	if err != nil {
		s.log.Errorln("Register :", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().WithMessage("Invalid data"))
	}

	err = validation.ValidateStruct(&payload,
		validation.Field(&payload.FullName, validation.Required, validation.Match(regexp.MustCompile(`^[A-Za-z\s]+$`)).
		Error("FullName only accept alphabet and space")),
		validation.Field(&payload.Email, validation.Required, is.Email),
		validation.Field(&payload.Password, validation.Required, validation.Length(6, 60), validation.NotIn(payload.Email).Error("Password cannot same with email")),
	)
	if err != nil {
		s.log.Errorln("Register :", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().WithMessage(err.Error()))
	}

	hashedPassword, err := pkg.GeneratePassword(payload.Password)
	if err != nil {
		s.log.Errorln("Register :", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().WithMessage("Invalid data"))
	}
	payload.Password = hashedPassword

	err = s.usersRepo.Register(c.Context(), payload)
	if err != nil {
		s.log.Errorln("Register :", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(resp.AsError().WithMessage("Internal server error"))
	}
	
	s.log.Infoln("Register success")
	return c.Status(fiber.StatusOK).JSON(resp.WithMessage("Register success"))
}