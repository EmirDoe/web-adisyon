package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"webadisyon.com/models"
)

func AddSingleTable(c *fiber.Ctx) error {
	table := new(models.Table)
	if err := c.BodyParser(table); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Get the latest tables number and increment it by one
	tables, err := models.GetTables()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	table.TableNumber = len(tables) + 1

	tableID, err := models.AddSingleTable(*table)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"table_id": tableID,
	})
}

func GetTables(c *fiber.Ctx) error {
	tables, err := models.GetTables()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"tables": tables,
	})
}

func AddTablesAtSetup(c *fiber.Ctx) error {
	// Get the number of tables to be added
	tableCount := c.Params("table_count")

	tableCountInt, _ := strconv.Atoi(tableCount)
	err := models.AddTablesAtSetup(tableCountInt)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Tables added successfully",
	})
	//ENV file tables setup will be 1 after this func -- Will be added
}
