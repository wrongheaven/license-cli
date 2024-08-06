package main

import (
	"fmt"
	"log"
	"net/mail"
	"os"
	"strings"
	"time"

	_ "embed"

	"github.com/charmbracelet/huh"
)

var (
	licenseType string
	authorName  string
	authorEmail string
)

//go:embed licenses/mit.md
var licenseMIT []byte

//go:embed licenses/apache.md
var licenseApache []byte

type License struct {
	License []byte
	Author  bool
}

var licenses map[string]License

func init() {
	licenses = make(map[string]License)
	licenses["mit"] = License{licenseMIT, true}
	licenses["apache"] = License{licenseApache, false}
}

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Which license to create?").
				Options(
					huh.NewOption("MIT", "mit"),
					huh.NewOption("Apache 2.0", "apache"),
				).
				Value(&licenseType),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	chosenLicense := licenses[licenseType]

	content := strings.ReplaceAll(
		string(chosenLicense.License),
		"{{YEAR}}",
		fmt.Sprint(time.Now().Year()),
	)

	if chosenLicense.Author {
		form = huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Author name").
					Value(&authorName),
				huh.NewInput().
					Title("Author email").
					Value(&authorEmail).
					Validate(func(str string) error {
						if _, err := mail.ParseAddress(str); err != nil {
							return err
						}
						return nil
					}),
			),
		)

		if err := form.Run(); err != nil {
			log.Fatal(err)
		}

		content = strings.ReplaceAll(
			string(content),
			"{{HOLDER}}",
			fmt.Sprintf("%s <%s>", authorName, authorEmail),
		)
	}

	err := os.WriteFile("LICENSE", []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("License created")
}
