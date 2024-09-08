package controllers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// This controller will be used at setup and afterwards change to the .env and it will change the values in .env by recreating it everytime and keeping the old values

// Base .env creation
func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf(".env couldnt be loaded: %v", err)
	}
	return nil
}

// GetEnvValue retrieves the value of a specific key from the .env file
func GetEnvValue(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("%s couldnt be found", key)
	}
	return value, nil
}

// UpdateEnvValue updates a specific key-value pair in the .env file
func UpdateEnvValue(key, newValue string) error {
	// .env dosyasını oku
	envMap, err := godotenv.Read(".env")
	if err != nil {
		return fmt.Errorf(".env couldnt be read: %v", err)
	}

	// Değerini güncelle
	envMap[key] = newValue

	// .env dosyasını yeniden yaz
	err = godotenv.Write(envMap, ".env")
	if err != nil {
		return fmt.Errorf(".env couldnt be written: %v", err)
	}

	return nil
}

// IsInstallationComplete checks if the INSTALLATION_DONE flag is set to true
func IsInstallationComplete() (bool, error) {
	value, err := GetEnvValue("INSTALLATION_DONE")
	if err != nil {
		return false, err
	}
	return value == "true", nil
}
