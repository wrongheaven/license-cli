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

func PCheck(err error) {
	if err != nil {
		panic(err)
	}
}

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

func GetLicencePath(licenseType string) string {
	exeDir, err := GetExeDir()
	PCheck(err)

	licenseTypeLower := strings.ToLower(licenseType)
	return filepath.Join(exeDir, ".license-cli", licenseTypeLower+".md")
}

func OpenTemplate(licenseType string) ([]byte, error) {
	licensePath := GetLicencePath(licenseType)
	content, err := os.ReadFile(licensePath)
	PCheck(err)

	return content, nil
}

func GetUserConfig() (models.User, error) {
	// Open ~/.config/license-cli/user.json
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return models.User{}, err
	}
	configFile, err := os.Open(filepath.Join(userHomeDir, ".config", "license-cli", "user.json"))
	if err != nil {
		return models.User{}, err
	}
	defer configFile.Close()

	byteValue, err := io.ReadAll(configFile)
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

func GetHomeDir() (string, error) {
	return os.UserHomeDir()
}
func GetExeDir() (string, error) {
	exe, err := os.Executable()
	PCheck(err)

	return filepath.Dir(exe), nil
}
