package main

import (
	"fmt"
	"os"
	"path/filepath"
	"postwp/validate"
)

func main() {
	programName := filepath.Base(os.Args[0])

	if len(os.Args) != 5 {
		fmt.Printf("Usage: %s WP_APP <app_name> DEV_DIRECTORY <path_name>\n", programName)
		return
	}

	if os.Args[1] != "WP_APP" || os.Args[3] != "DEV_DIRECTORY" {
		fmt.Printf("Usage: %s WP_APP <app_name> DEV_DIRECTORY <path_name>\n", programName)
		return
	}
	wpApp := os.Args[2]
	devDirectory := os.Args[4]

	fmt.Printf("WP_APP: %s\n", wpApp)
	fmt.Printf("DEV_DIRECTORY: %s\n", devDirectory)

	validate.EnsureDirectory(devDirectory)
}
