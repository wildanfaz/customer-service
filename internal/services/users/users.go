package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/go-market/configs"
	"github.com/wildanfaz/go-market/internal/repositories"
)

type ImplementServices struct {
	log *logrus.Logger
	usersRepo repositories.Users
	config *configs.Config
}

type Services interface{
	Register(c *fiber.Ctx) error
	Login (c *fiber.Ctx) error
}

func New(log *logrus.Logger, usersRepo repositories.Users, config *configs.Config) Services {
	return &ImplementServices{
		log: log,
		usersRepo: usersRepo,
		config: config,
	}
}