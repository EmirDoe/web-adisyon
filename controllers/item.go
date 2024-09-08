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

func UpdateItemCategory(c *fiber.Ctx) error {
	request := &ItemCategoryRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	categoryID, err := c.ParamsInt("category_id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	record := models.ItemCategory{
		Name: request.Name,
	}

	err = models.UpdateItemCategory(categoryID, record)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(204)
}

func DeleteItemCategory(c *fiber.Ctx) error {
	categoryID, err := c.ParamsInt("category_id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = models.DeleteItemCategory(categoryID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(204)
}
func GetItemCategory(c *fiber.Ctx) error {
	categoryID, err := c.ParamsInt("category_id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	category, err := models.GetItemCategory(categoryID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(category)
}

func GetItemCategories(c *fiber.Ctx) error {
	categories, err := models.GetItemCategories()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(categories)
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

func UpdateItem(c *fiber.Ctx) error {
	request := &ItemRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	itemID := c.Params("item_id")

	record := models.Item{
		MenuID:      request.MenuID,
		Category:    request.Category,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		ImagePath:   request.ImagePath,
	}

	err := models.UpdateItem(itemID, record)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(204)
}

func DeleteItem(c *fiber.Ctx) error {
	itemID := c.Params("item_id")

	err := models.DeleteItem(itemID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(204)
}

func GetItems(c *fiber.Ctx) error {
	items, err := models.GetItems()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(items)
}

func GetItemByID(c *fiber.Ctx) error {
	itemID := c.Params("item_id")

	item, err := models.GetItemByID(itemID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(item)
}
