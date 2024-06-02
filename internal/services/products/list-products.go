package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanfaz/go-market/internal/helpers"
	"github.com/wildanfaz/go-market/internal/models"
)

func (s *ImplementServices) ListProducts(c *fiber.Ctx) error {
	var (
		resp = helpers.NewResponse()
		payload models.Product
	)

	err := c.QueryParser(&payload)
	if err != nil {
		s.log.Errorln("ListProducts :", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().WithMessage("Invalid data"))
	}

	products, err := s.productsRepo.ListProducts(c.Context(), payload)
	if err != nil {
		s.log.Errorln("ListProducts :", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(resp.AsError().WithMessage("Internal server error"))
	}

	s.log.Infoln("List products success")
	return c.Status(fiber.StatusOK).JSON(resp.WithMessage("List products success").WithData(products))
}