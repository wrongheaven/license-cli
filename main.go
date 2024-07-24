package main

import (
	"fmt"
	"os"
	"wrongheaven/license-cli/license"
	"wrongheaven/license-cli/utils"
)

func main() {
	arglen := len(os.Args)

	if arglen == 1 {
		utils.ShowUsage("main")
	}

	switch command := os.Args[1]; command {
	case "config":
		fmt.Println("<config not implemented>")
	case "add":
		if len(os.Args) == 2 {
			utils.ShowUsage("add")
		}

		licenseType := os.Args[2]
		license.Add(licenseType)
	case "show":
		fmt.Println("<show not implemented>")
	default:
		utils.ShowUsage("main")
	}
}
