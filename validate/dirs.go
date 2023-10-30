package validate

import (
	"fmt"
	"os"
	"path/filepath"
)

func EnsureDirectory(devDirectory string) bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error retrieving home directory:", err)
		return false
	}

	dirNameWithSuffix := devDirectory + "-wp-svc"
	dirToCheck := filepath.Join(homeDir, "postwp", dirNameWithSuffix)

	if _, err := os.Stat(dirToCheck); os.IsNotExist(err) {
		fmt.Printf("Directory '%s' does not exist. Creating...\n", dirToCheck)
		if mkdirErr := os.MkdirAll(dirToCheck, 0755); mkdirErr != nil {
			fmt.Printf("Failed to create directory '%s'. Error: %s\n", dirToCheck, mkdirErr)
			return false
		}
		fmt.Printf("Directory '%s' created successfully.\n", dirToCheck)
	} else {
		fmt.Printf("Directory '%s' exists.\n", dirToCheck)
	}
	return true
}
