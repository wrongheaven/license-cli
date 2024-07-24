package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"wrongheaven/license-cli/models"
)

func IsGoRun() bool {
	return strings.Contains(os.Args[0], "go-build")
}

func ShowUsage(command string) {
	switch command {
	case "main":
		fmt.Println("Usage: license-cli [add|show|config] [...params]")
	case "add":
		fmt.Println("Usage: license-cli add [mit|apache]")
	case "show":
		fmt.Println("Usage: license-cli show [mit|apache]")
	}

	os.Exit(0)
}

func GetLicensePath(licenseType string) (string, error) {
	if IsGoRun() {
		return filepath.Join("licenses", licenseType+".md"), nil
	} else {
		ex, err := os.Executable()
		if err != nil {
			return "", err
		}
		return filepath.Join(filepath.Dir(ex), "licenses", licenseType+".md"), nil
	}
}

func GetUserConfig() (models.User, error) {
	// Open ~/.config/license-cli
	_, err := os.UserHomeDir()
	if err != nil {
		return models.User{}, err
	}
	// configFile, err := os.Open(filepath.Join(userHomeDir, ".config", "license-cli", "user.json"))
	config, err := os.Open(filepath.Join("config", "user.json"))
	if err != nil {
		return models.User{}, err
	}
	defer config.Close()

	byteValue, err := io.ReadAll(config)
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = json.Unmarshal(byteValue, &user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
