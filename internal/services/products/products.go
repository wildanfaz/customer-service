package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/go-market/configs"
	"github.com/wildanfaz/go-market/internal/repositories"
)

type ImplementServices struct {
	log *logrus.Logger
	productsRepo repositories.Products
	config *configs.Config
}

type Services interface{
	ListProducts(c *fiber.Ctx) error
}

func New(log *logrus.Logger, productsRepo repositories.Products, config *configs.Config) Services {
	return &ImplementServices{
		log: log,
		productsRepo: productsRepo,
		config: config,
	}
}