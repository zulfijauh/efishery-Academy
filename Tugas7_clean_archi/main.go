package main

import (
	"clean-archi/config"
	"clean-archi/handler"
	"clean-archi/repository"
	"clean-archi/routes"
	"clean-archi/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Database()
	config.Migrate()

	e := echo.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	UserHandler := handler.NewUserHandler(userUsecase)

	routes.Routes(e, UserHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
