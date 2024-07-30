package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "embed"
)

//go:embed templates/mit.md
var licenseMIT []byte

/*
Fatally exits if error exists.
Please be nice.
*/
func benice(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// func yesOrNo(licenseType *string, msg string) error {
// 	fmt.Printf("%s (Y/n) ", msg)
// 	var resp string
// 	fmt.Scanln(&resp)
// 	if resp == "" || strings.ToLower(resp) == "y" {
// 		return nil
// 	} else {
// 		return errors.New("`no` selected")
// 	}
// }

func main() {
	const licenseType string = "mit"
	const authorName string = "Henrik Engebretsen"
	const authorEmail string = "wrongheaven73@gmail.com"

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

	// 2. save new license to CWD
	err := os.WriteFile("LICENSE", []byte(newContent), 0666)
	benice(err)

	// fmt.Println("License created:")
	// fmt.Println("")
	// fmt.Println(newContent)

	fmt.Printf("License created:\n\n%s", newContent)
}
