package main

import (
	"assessment_efishery/config"
	"assessment_efishery/handler"
	"assessment_efishery/repository"
	"assessment_efishery/routes"
	"assessment_efishery/usecase"

	"github.com/labstack/echo/v4"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
func main() {
	config.Database()
	config.Migrate()

	e := echo.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	UserHandler := handler.NewUserHandler(userUsecase)
	routes.UserRoutes(e, UserHandler)

	productRepository := repository.NewProductRepository(config.DB)
	productUsecase := usecase.NewProductUsecase(productRepository)
	ProductHandler := handler.NewProductHandler(productUsecase)
	routes.ProductRoutes(e, ProductHandler)

	cartRepository := repository.NewCartRepository(config.DB)
	cartUsecase := usecase.NewCartUsecase(cartRepository)
	CartHandler := handler.NewCartHandler(cartUsecase)
	routes.CartRoutes(e, CartHandler)

	transactionRepository := repository.NewTransactionRepository(config.DB)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepository)
	TransactionHandler := handler.NewTransactionHandler(transactionUsecase)
	routes.TransactionRoutes(e, TransactionHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
