package controllers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"webadisyon.com/models"
)

type AssetUploadRequest struct {
	ItemID string `json:"item_id"`
}

func UploadImage(c *fiber.Ctx) error {

	//Validate user
	userAuth := ValidateUser(c)
	if !userAuth.IsAuthenticated {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// parse request body

	// Parse the multipart form file
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	assetUploadRequest := form.Value["item_id"][0]
	// Get the files from form
	files := form.File["image"]
	if len(files) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "No file uploaded",
		})
	}

	// Get the first file
	file := files[0]

	// Move the file to the assets directory
	// Create the assets directory if it doesn't exists
	assetsDir := filepath.Join("assets")
	if _, err := os.Stat(assetsDir); os.IsNotExist(err) {
		os.Mkdir(assetsDir, 0755)
	}

	file.Filename = uuid.New().String() + filepath.Ext(file.Filename)
	filExt := filepath.Ext(file.Filename)

	// Move the file to the assets directory
	if err := c.SaveFile(file, filepath.Join(assetsDir, file.Filename)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// remove extension from filename
	file.Filename = strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))

	if err := models.UploadItemAsset(models.Asset{
		ID:   file.Filename,
		Name: file.Filename,
		Path: filepath.Join(assetsDir, file.Filename+filExt),
	}, assetUploadRequest); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Image uploaded",
	})

}
