package carts

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/wildanfaz/go-market/internal/helpers"
	"github.com/wildanfaz/go-market/internal/models"
)

func (s *ImplementServices) AddToCart(c *fiber.Ctx) error {
	var (
		resp    = helpers.NewResponse()
		payload models.Cart
	)

	user, ok := c.Locals("user").(*models.User)
	if !ok {
		s.log.Errorln("AddToCart : error parsing user")
		return c.Status(fiber.StatusInternalServerError).JSON(resp.AsError().WithMessage("Internal server error"))
	}

	err := c.BodyParser(&payload)
	if err != nil {
		s.log.Errorln("AddToCart :", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().WithMessage("Invalid data"))
	}

	err = validation.ValidateStruct(&payload,
		validation.Field(&payload.ProductId, validation.Required, validation.Min(1)),
		validation.Field(&payload.Quantity, validation.Required, validation.Min(1)),
	)
	if err != nil {
		s.log.Errorln("AddToCart :", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().WithMessage(err.Error()))
	}

	payload.UserId = user.Id
	
	err = s.cartsRepo.AddToCart(c.Context(), payload)
	if err != nil {
		s.log.Errorln("AddToCart :", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(resp.AsError().WithMessage("Internal server error"))
	}

	s.log.Infoln("Add to cart success")
	return c.Status(fiber.StatusOK).JSON(resp.WithMessage("Add to cart success"))
}