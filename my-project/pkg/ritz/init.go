package ritz

import (
	"fmt"
	"os"
	"path/filepath"

	"my-project/pkg/utility"
)

func Init() {
	if utility.IsInitialized() {
		fmt.Println("The directory is already initialized as a ritz directory.")
		return
	}

	// Create the .ritz directory structure
	if err := os.Mkdir(".ritz", 0755); err != nil {
		fmt.Println("Failed to create .ritz directory:", err)
		return
	}

	// Create objects directory
	objectsPath := filepath.Join(".ritz", "objects")
	if err := os.Mkdir(objectsPath, 0755); err != nil {
		fmt.Println("Failed to create objects directory:", err)
		return
	}

	// Create refs/heads directory
	refsPath := filepath.Join(".ritz", "refs", "heads")
	if err := os.MkdirAll(refsPath, 0755); err != nil {
		fmt.Println("Failed to create refs/heads directory:", err)
		return
	}

	// Create HEAD file
	headContent := "ref: refs/heads/master\n"
	if err := os.WriteFile(filepath.Join(".ritz", "HEAD"), []byte(headContent), 0644); err != nil {
		fmt.Println("Failed to create HEAD file:", err)
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current directory:", err)
		return
	}

	fmt.Println("Initialized empty ritz directory in", filepath.Join(currentDir, ".ritz"))
}

