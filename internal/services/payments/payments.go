package payments

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/go-market/internal/repositories"
)

type ImplementServices struct {
	log *logrus.Logger
	paymentsRepo repositories.Payments
}

type Services interface{
	Pay(c *fiber.Ctx) error
}

func New(log *logrus.Logger, paymentsRepo repositories.Payments) Services {
	return &ImplementServices{
		log: log,
		paymentsRepo: paymentsRepo,
	}
}