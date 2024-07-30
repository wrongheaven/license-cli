package utils

import (
	"fmt"
	"os"
	"path/filepath"

	_ "embed"
)

func GetLicense(licenseType string) ([]byte, error) {

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	licensePath := filepath.Join(userHomeDir, ".local", "license-cli")

	fmt.Println(licensePath)

	licenseFile, err := os.ReadFile(licensePath)
	if err != nil {
		return nil, err
	}

	return licenseFile, nil
}

func AddLicense(licenseType string) error {
	// Get license from ~/.local/license-cli/templates/[licenseType].md
	a, err := GetLicense(licenseType)
	if err != nil {
		return err
	}

	fmt.Println("Add(): ", a)

	// Get name and email from config

	// Replace placeholders in license with those from config

	// Write new license to some folder

	//
	//
	//
	//
	//
	//

	return nil

	// fmt.Println(licenseType + "1")

	// licenseFile, err := utils.GetLicensePath(licenseType)
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(licenseType + "2")

	// // Open license template for reading
	// content, err := os.ReadFile(licensePath)
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(licenseType + "3")

	// // Get name and email from config
	// user, err := utils.GetUserConfig()
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(licenseType + "4")

	// // Replace placeholders with corresponding values
	// lines := strings.Split(string(content), "\n")
	// for i, line := range lines {
	// 	lines[i] = strings.ReplaceAll(line, "{{YEAR}}", "2024")
	// 	lines[i] = strings.ReplaceAll(
	// 		lines[i],
	// 		"{{HOLDER}}",
	// 		fmt.Sprintf("%s <%s>", user.Name, user.Email),
	// 	)
	// }

	// fmt.Println(licenseType + "5")

	// // Construct license file from updated lines
	// newContent := strings.Join(lines, "\n")
	// err = os.WriteFile("LICENSE", []byte(newContent), 0644)
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(licenseType + "6")

	// return nil
}
