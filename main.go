package main

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"webadisyon.com/db"
	"webadisyon.com/routes"
)

func init() {
	db.Setup()
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		//AllowCredentials: true,
	}))

	app.Use(InstallationCheck)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.Routes(app)
	app.Listen(":8080")

}

func InstallationCheck(c *fiber.Ctx) error {
	installationDone := os.Getenv("INSTALLATION_DONE")

	if installationDone != "true" {

		// if path does not start with /inst

		if !strings.HasPrefix(c.Path(), "/api/v1/install") {
			return c.Status(403).JSON(fiber.Map{
				"message":  "Installation isnt done. Please finish installation.",
				"redirect": "/install",
			})
		}
	}

	return c.Next()
}
