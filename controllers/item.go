package controllers

import (
	"github.com/gofiber/fiber/v2"
	"webadisyon.com/models"
)

type ItemRequest struct {
	MenuID      int    `json:"menu_id"`
	Category    int    `json:"category"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImagePath   string `json:"image"`
}

type ItemCategoryRequest struct {
	Name string `json:"name"`
}

func AddItemCategory(c *fiber.Ctx) error {
	request := &ItemCategoryRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	record := models.ItemCategory{
		Name: request.Name,
	}

	// USER VALIDATION WIL BE HERE

	categoryID, err := models.AddItemCategory(record)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"category_id": categoryID,
	})

}

func AddItem(c *fiber.Ctx) error {
	request := &ItemRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// USER VALIDATION WIL BE HERE

	record := models.Item{
		MenuID:      request.MenuID,
		Category:    request.Category,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		ImagePath:   request.ImagePath,
	}

	itemID, err := models.AddItem(record)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"item_id": itemID,
	})

}
