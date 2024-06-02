package payments

import (
	"github.com/gofiber/fiber/v2"

	"github.com/wildanfaz/go-market/internal/constants"
	"github.com/wildanfaz/go-market/internal/helpers"
	"github.com/wildanfaz/go-market/internal/models"
)

func (s *ImplementServices) Pay(c *fiber.Ctx) error {
	var (
		resp = helpers.NewResponse()
	)

	user, ok := c.Locals("user").(*models.User)
	if !ok {
		s.log.Errorln("Pay : error parsing user")
		return c.Status(fiber.StatusInternalServerError).JSON(resp.AsError().WithMessage("Internal server error"))
	}

	err := s.paymentsRepo.Pay(c.Context(), *user)
	if err != nil && err.Error() == constants.InsufficientBalance {
		s.log.Errorln("Pay :", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().WithMessage(constants.InsufficientBalance))
	}

	if err != nil {
		s.log.Errorln("Pay :", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(resp.AsError().WithMessage("Internal server error"))
	}

	s.log.Infoln("Pay success")
	return c.Status(fiber.StatusOK).JSON(resp.WithMessage("Pay success"))
}