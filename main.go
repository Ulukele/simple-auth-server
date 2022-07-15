package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	server, err := CreateServer()
	if err != nil {
		log.Fatal(err)
	}

	authGroup := app.Group("/api/auth/")
	authGroup.Get("/sign-in/", server.HandleAuthSignIn)
	authGroup.Get("/validate/", server.HandleAuthValidate)

	log.Fatal(app.Listen(":8080"))
}
