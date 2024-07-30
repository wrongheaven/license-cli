package main

import (
	"fmt"
	"log"
	"strings"

	_ "embed"
)

//go:embed templates/mit.md
var licenseMIT []byte

func benice(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	const licenseType string = "mit"
	const authorName string = "John Doe"
	const authorEmail string = "john@doe.com"

	// 1. read the license from file
	var content []byte
	switch strings.ToLower(licenseType) {
	case "mit":
		content = licenseMIT
	default:
		log.Fatal("invalid license")
	}

	// 1a. replace placeholders with info
	newContent := strings.ReplaceAll(string(content), "{{YEAR}}", "2024")
	newContent = strings.ReplaceAll(
		string(newContent),
		"{{HOLDER}}",
		fmt.Sprintf("%s <%s>", authorName, authorEmail),
	)

	fmt.Println(newContent)

	// 2. save new license to CWD
}
