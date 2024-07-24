package license

import (
	"fmt"
	"os"
	"strings"

	"wrongheaven/license-cli/utils"
)

func Add(licenseType string) error {

	licensePath, err := utils.GetLicensePath(licenseType)
	if err != nil {
		return err
	}

	// Open license template for reading
	content, err := os.ReadFile(licensePath)
	if err != nil {
		return err
	}

	// Get name and email from config
	user, err := utils.GetUserConfig()
	if err != nil {
		return err
	}

	// Replace placeholders with corresponding values
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, "{{YEAR}}", "2024")
		lines[i] = strings.ReplaceAll(
			lines[i],
			"{{HOLDER}}",
			fmt.Sprintf("%s <%s>", user.Name, user.Email),
		)
	}

	// Construct license file from updated lines
	newContent := strings.Join(lines, "\n")
	err = os.WriteFile("LICENSE", []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
