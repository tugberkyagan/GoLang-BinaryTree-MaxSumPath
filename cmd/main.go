package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-wrk/pkg/handlers"
	"go-wrk/pkg/services"
	"go-wrk/utils"
	"log"
)

func main() {

	config := utils.NewConfiguration()

	app := fiber.New()

	binaryTreeService := services.BinaryTreeServiceStruct{}

	binaryTreeHandler := handlers.BinaryTreeHandlerStruct{
		Service: binaryTreeService,
	}

	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	app.Post("/calculate-max-path-sum", binaryTreeHandler.CalculateMaxPathSum)

	err := app.Listen(config.Server.Port)

	if err != nil {
		log.Println("Fiber listening problem", err)
	}
}
