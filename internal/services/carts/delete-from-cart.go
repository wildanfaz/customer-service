package carts

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/wildanfaz/go-market/internal/helpers"
	"github.com/wildanfaz/go-market/internal/models"
)

func (s *ImplementServices) DeleteFromCart(c *fiber.Ctx) error {
	var (
		resp = helpers.NewResponse()
		id = c.Params("id")
	)

	user, ok := c.Locals("user").(*models.User)
	if !ok {
		s.log.Errorln("DeleteFromCart : error parsing user")
		return c.Status(fiber.StatusInternalServerError).JSON(resp.AsError().WithMessage("Internal server error"))
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		s.log.Errorln("DeleteFromCart : error parsing id")
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().WithMessage("Invalid data"))
	}

	if idInt < 1 {
		s.log.Errorln("DeleteFromCart : Id cannot be less than 1")
		return c.Status(fiber.StatusBadRequest).JSON(resp.AsError().WithMessage("Id cannot be less than 1"))
	}

	err = s.cartsRepo.DeleteFromCart(c.Context(), user.Id, idInt)
	if err != nil && err == sql.ErrNoRows {
		s.log.Errorln("DeleteFromCart :", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(resp.AsError().WithMessage("Data not found"))
	}

	if err != nil {
		s.log.Errorln("DeleteFromCart :", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(resp.AsError().WithMessage("Internal server error"))
	}

	s.log.Infoln("Delete from cart success")
	return c.Status(fiber.StatusOK).JSON(resp.WithMessage("Delete from cart success"))
}