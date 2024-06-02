package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanfaz/go-market/configs"
	"github.com/wildanfaz/go-market/internal/middlewares"
	"github.com/wildanfaz/go-market/internal/pkg"
	"github.com/wildanfaz/go-market/internal/repositories"
	"github.com/wildanfaz/go-market/internal/services/carts"
	"github.com/wildanfaz/go-market/internal/services/payments"
	"github.com/wildanfaz/go-market/internal/services/products"
	"github.com/wildanfaz/go-market/internal/services/users"
)

func InitRouter() {
	// configs
	config := configs.InitConfig()
	db := configs.InitMySql(config.DatabaseDSN)

	// pkg
	log := pkg.InitLogger()

	// repositories
	usersRepo := repositories.NewUsersRepository(db)
	productsRepo := repositories.NewProductsRepository(db)
	cartsRepo := repositories.NewCartsRepository(db)
	paymentsRepo := repositories.NewPaymentsRepository(db)

	// services
	usersServices := users.New(log, usersRepo, config)
	productsServices := products.New(log, productsRepo, config)
	cartsServices := carts.New(log, cartsRepo)
	paymentsServices := payments.New(log, paymentsRepo)

	// middlewares
	auth := middlewares.Auth(log, config, usersRepo)

	app := fiber.New()

	apiV1 := app.Group("/api/v1")

	// users
	usersV1 := apiV1.Group("/users")
	usersV1.Post("/register", usersServices.Register)
	usersV1.Post("/login", usersServices.Login)

	apiV1.Use(auth)

	// products
	productsV1 := apiV1.Group("/products")
	productsV1.Get("/list-products", productsServices.ListProducts)

	// carts
	cartsV1 := apiV1.Group("/carts")
	cartsV1.Post("/add-to-cart", cartsServices.AddToCart)
	cartsV1.Get("/list-in-cart", cartsServices.ListInCart)
	cartsV1.Delete("/delete-from-cart/:id", cartsServices.DeleteFromCart)

	// payments
	paymentsV1 := apiV1.Group("/payments")
	paymentsV1.Post("/pay", paymentsServices.Pay)

	app.Listen(config.AppPort)
}
