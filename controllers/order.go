package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"webadisyon.com/models"
)

type OrderRequest struct {
	TableID   string               `json:"table_id"`
	Items     []models.ItemOnOrder `json:"items"`
	Total     int                  `json:"total"`
	Status    string               `json:"status"`
	CreatedBy string               `json:"created_by"`
	CreatedAt string               `json:"created_at"`
}

type ActionRequest struct {
	OrderID    string               `json:"order_id"`
	UserID     string               `json:"user_id"`
	Items      []models.ItemOnOrder `json:"items"`
	ActionType string               `json:"action_type"`
	Timestamp  time.Time            `json:"timestamp"`
}

func CreateOrder(c *fiber.Ctx) error {
	request := &OrderRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Check if the request is valid
	if request.TableID == "" || len(request.Items) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	record := models.Order{
		TableID:   request.TableID,
		Items:     request.Items,
		Total:     request.Total,
		Status:    "OPEN",
		CreatedBy: request.CreatedBy,
		CreatedAt: request.CreatedAt,
	}
	record.Total = models.CalculateTotal(record)

	userAuth := ValidateUser(c)
	if !userAuth.IsAuthenticated {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	record.Actions = append(record.Actions, models.Action{
		UserID:     userAuth.UserID,
		Items:      request.Items,
		ActionType: "CREATE",
		Timestamp:  time.Now(),
	})

	orderID, err := models.CreateOrder(record)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"order_id": orderID,
	})
}

func AddAction(c *fiber.Ctx) error {
	request := &ActionRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	record := models.Action{
		UserID:     request.UserID,
		Items:      request.Items,
		ActionType: "ADD",
		Timestamp:  time.Now(),
	}

	userAuth := ValidateUser(c)
	if !userAuth.IsAuthenticated {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	err := models.AddAction(request.OrderID, record)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "ADD Action added successfully",
	})
}

func RemoveAction(c *fiber.Ctx) error {
	request := &ActionRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	record := models.Action{
		UserID:     request.UserID,
		Items:      request.Items,
		ActionType: "REMOVE",
		Timestamp:  time.Now(),
	}

	userAuth := ValidateUser(c)
	if !userAuth.IsAuthenticated {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	err := models.RemoveAction(request.OrderID, record)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "REMOVE Action added successfully",
	})
}
