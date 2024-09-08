package controllers

import (
	"github.com/gofiber/fiber/v2"

	"webadisyon.com/db"
)

type RestaurantInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}

type DBConfig struct {
	DBUri  string `json:"db_uri"`
	DBName string `json:"db_name"`
}

type AdminUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Restaurant info setup

func SetupRestaurantInfo(c *fiber.Ctx) error {
	restaurantInfo := &RestaurantInfo{}
	if err := c.BodyParser(restaurantInfo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if restaurantInfo.Name == "" || restaurantInfo.Description == "" || restaurantInfo.Location == "" || restaurantInfo.Phone == "" || restaurantInfo.Email == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	UpdateEnvValue("RESTAURANT_NAME", restaurantInfo.Name)
	UpdateEnvValue("RESTAURANT_DESCRIPTION", restaurantInfo.Description)
	UpdateEnvValue("RESTAURANT_LOCATION", restaurantInfo.Location)
	UpdateEnvValue("RESTAURANT_PHONE", restaurantInfo.Phone)
	UpdateEnvValue("RESTAURANT_EMAIL", restaurantInfo.Email)

	return c.JSON(fiber.Map{
		"message": "Restaurant info saved",
	})
}

func SetupDBConfig(c *fiber.Ctx) error {
	dbConfig := &DBConfig{}
	if err := c.BodyParser(dbConfig); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if dbConfig.DBUri == "" || dbConfig.DBName == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	UpdateEnvValue("MONGO_URI", dbConfig.DBUri)
	UpdateEnvValue("MONGO_DB_NAME", dbConfig.DBName)

	//Test the connection
	log := db.Client.Database(dbConfig.DBName).Client().Ping(db.Context, nil)
	if log != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": log.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "DB config saved",
	})
}
