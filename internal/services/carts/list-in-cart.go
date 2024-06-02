package carts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanfaz/go-market/internal/helpers"
	"github.com/wildanfaz/go-market/internal/models"
)

func (s *ImplementServices) ListInCart(c *fiber.Ctx) error {
	var (
		resp = helpers.NewResponse()
	)

	user, ok := c.Locals("user").(*models.User)
	if !ok {
		s.log.Errorln("ListInCart : error parsing user")
		return c.Status(fiber.StatusInternalServerError).JSON(resp.AsError().WithMessage("Internal server error"))
	}

	carts, err := s.cartsRepo.ListInCart(c.Context(), user.Id)
	if err != nil {
		s.log.Errorln("ListInCart :", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(resp.AsError().WithMessage("Internal server error"))
	}

	s.log.Infoln("List in cart success")
	return c.Status(fiber.StatusOK).JSON(resp.WithMessage("List in cart success").WithData(carts))
}