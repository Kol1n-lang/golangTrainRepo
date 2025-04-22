package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"train-http/config"
	"train-http/internal/handlers"
	"train-http/internal/middleware"
	"train-http/internal/validators"
	"train-http/pkg/database"
)

func initServer() *echo.Echo {
	e := echo.New()
	e.Validator = &validators.CustomValidator{Validator: validator.New()}
	return e
}

func initHandlers(e *echo.Echo) {
	e.POST("register", handlers.RegisterUser)
	restricted := e.Group("")
	restricted.Use(middleware.CheckAccess)
	restricted.GET("/get-user", handlers.GetUser)
	restricted.GET("/docs", handlers.DocsInfo)
}

func main() {
	cfg, _ := config.LoadConfig("config.toml")
	server := initServer()
	initHandlers(server)
	database.InitDB()

	err := server.Start(cfg.Server_URL())
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
