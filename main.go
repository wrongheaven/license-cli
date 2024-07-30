package main

import (
	"fmt"
	"os"
	"path/filepath"
	"wrongheaven/license-cli/utils"
)

func main() {

	home, err := utils.GetHomeDir()
	utils.PCheck(err)
	exe, err := utils.GetExeDir()
	utils.PCheck(err)

	fmt.Println("home dir:", home)
	fmt.Println("exe dir:", exe)

	os.Exit(0)

	//

	// Open MIT template
	template, err := utils.OpenTemplate("mit")
	utils.PCheck(err)

	fmt.Println(template)

	os.Exit(0)

	homeDir, err := os.UserHomeDir()
	utils.PCheck(err)
	_, _ = os.ReadFile(filepath.Join(homeDir))

	exeDir, err := utils.GetExeDir()
	utils.PCheck(err)
	fmt.Println("exedir:", exeDir)

	os.Exit(0)

	// Does $HOME/.license-cli/ exist?
	rootDirExists := dirExists(".")
	fmt.Println(rootDirExists)
	os.Exit(0)

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
		err := utils.AddLicense(licenseType)
		if err != nil {
			panic(err)
		}
	case "show":
		fmt.Println("<show not implemented>")
	default:
		utils.ShowUsage("main")
	}
}

func dirExists(path string) bool {
	homeDir, err := os.UserHomeDir()
	utils.PCheck(err)
	p := filepath.Join(homeDir, ".license-cli", path)
	fmt.Println("p:", p)

	asd, err := os.Stat(filepath.Join(homeDir, ".license-cli", path))
	utils.PCheck(err)

	return asd.IsDir()
}
