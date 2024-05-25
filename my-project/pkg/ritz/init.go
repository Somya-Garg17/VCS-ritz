package main

import (
	"fmt"
	"os"
)

func main() {

	if _, err := os.Stat(".ritz"); err == nil {
		fmt.Println("The directory is already a ritz directory")
		return
	}

	if err := os.Mkdir(".ritz", 0755); err != nil {
		fmt.Println("Cannot create a .ritz directory:", err)
		return
	}

	if err := os.Mkdir(".ritz/objects", 0755); err != nil {
		fmt.Println("Cannot create objects directory:", err)
		return

	}

	if err := os.Mkdir(".ritz/refs", 0755); err != nil {
		fmt.Println("Cannot create refs directory:", err)
		return
	}

	headContent := []byte("ref: refs/heads/master\n")
	if err := os.WriteFile(".ritz/HEAD", headContent, 0644); err != nil {
		fmt.Println("Cannot create HEAD file:", err)
		return
	}

	fmt.Println("An empty ritz repository has been initialized")
}
