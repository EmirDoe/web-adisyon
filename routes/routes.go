package routes

import (
	"webadisyon.com/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(a *fiber.App) {
	route := a.Group("/api/v1")

	// Installation
	route.Post("/install/restaurant-info", controllers.SetupRestaurantInfo)
	route.Post("/install/db", controllers.SetupDBConfig)

	// Auth
	route.Post("/register", controllers.Register)
	route.Post("/login", controllers.Login)

	// Tables
	route.Get("/tables", controllers.GetTables)
	route.Post("/setup/tables/:table_count", controllers.AddTablesAtSetup)
	route.Post("/table", controllers.AddSingleTable)

	// Items
	route.Post("/item", controllers.AddItem)
	route.Get("/items", controllers.GetItems)
	route.Get("/item/:item_id", controllers.GetItemByID)
	route.Put("/item/:item_id", controllers.UpdateItem)
	route.Delete("/item/:item_id", controllers.DeleteItem)

	route.Post("/item/category", controllers.AddItemCategory)
	route.Get("/item/categories", controllers.GetItemCategories)
	route.Get("/item/category/:category_id", controllers.GetItemCategory)
	route.Put("/item/category/:category_id", controllers.UpdateItemCategory)
	route.Delete("/item/category/:category_id", controllers.DeleteItemCategory)

	// Orders
	route.Post("/order", controllers.CreateOrder)
	route.Post("/order/action", controllers.AddAction)
	route.Delete("/order/action", controllers.RemoveAction)

	// Assets
	route.Post("/upload/image", controllers.UploadImage)

}
