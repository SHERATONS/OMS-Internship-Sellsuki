package FiberServer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ConnectFiberServer() {
	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	app := fiber.New()

	app.Get("/products", func(c *fiber.Ctx) error {
		return c.JSON(GetProducts)
	})

	if err := app.Listen(port); err != nil {
		log.Fatal(err)
	}
}
