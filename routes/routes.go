package routes

import (
	"webadisyon.com/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(a *fiber.App) {
	route := a.Group("/api/v1")

	// Tables
	route.Get("/tables", controllers.GetTables)
	route.Post("/setup/tables/:table_count", controllers.AddTablesAtSetup)
	route.Post("/table", controllers.AddSingleTable)

	// Items
	route.Post("/item/category", controllers.AddItemCategory)
	route.Post("/item", controllers.AddItem)

}
