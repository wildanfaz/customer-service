package carts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/go-market/internal/repositories"
)

type ImplementServices struct {
	log *logrus.Logger
	cartsRepo repositories.Carts
}

type Services interface {
	AddToCart(c *fiber.Ctx) error
	ListInCart(c *fiber.Ctx) error
	DeleteFromCart(c *fiber.Ctx) error
}

func New(log *logrus.Logger, cartsRepo repositories.Carts) Services {
	return &ImplementServices{
		log:    log,
		cartsRepo: cartsRepo,
	}
}