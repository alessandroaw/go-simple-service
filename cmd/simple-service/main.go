package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/therealsandro/go-simple-service/pkg/users"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userRepo := users.NewUserRepository()
	uuc := users.NewUserUseCase(userRepo)
	users.NewUserHandler(e, uuc)

	// Start Server
	e.Logger.Fatal(e.Start(":4000"))
}
