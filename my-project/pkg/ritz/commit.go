package ritz

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"my-project/pkg/utility"
)

func Commit(message string) {
	if !utility.IsInitialized() {
		fmt.Println("This is not a ritz repository. Run 'ritz init' to initialize one.")
		return
	}

	indexPath := filepath.Join(".ritz", "index")
	index := utility.LoadIndex(indexPath)

	// Create commit tree object
	treeHash, err := createTreeObject(index)
	if err != nil {
		fmt.Println("Error creating tree object:", err)
		return
	}

	// Write commit object
	commitHash, err := createCommitObject(treeHash, message)
	if err != nil {
		fmt.Println("Error writing commit object:", err)
		return
	}

	// Update HEAD reference
	if err := updateRef("refs/heads/master", commitHash); err != nil {
		fmt.Println("Error updating HEAD reference:", err)
		return
	}

	// Get the list of files from the index
	files := getFilesFromIndex(index)
	
	// Get the number of files added or changed
	numFiles := len(files)

	fmt.Printf("[%s (root-commit) %s] %s\n", "master", commitHash, message)
	fmt.Printf(" %d file(s) changed\n", numFiles)
	for _, file := range files {
		fmt.Printf(" create mode 100644 %s\n", file)
	}
}

func createTreeObject(index map[string]string) (string, error) {
	var lines []string
	for file, hash := range index {
		lines = append(lines, fmt.Sprintf("100644 %s\t%s", hash, file))
	}
	treeContent := strings.Join(lines, "\n")

	treeHash := utility.Sha1([]byte(treeContent))
	treePath := filepath.Join(".ritz", "objects", treeHash[:2])
	if err := os.MkdirAll(treePath, 0755); err != nil {
		return "", err
	}
	treeFilePath := filepath.Join(treePath, treeHash[2:])
	if err := utility.WriteObject(treeFilePath, []byte(treeContent)); err != nil {
		return "", err
	}
	return treeHash, nil
}

func createCommitObject(treeHash, message string) (string, error) {
	commitContent := fmt.Sprintf("tree %s\n", treeHash)
	commitContent += "author Your Name <you@example.com> 0 +0000\n"
	commitContent += "committer Your Name <you@example.com> 0 +0000\n\n"
	commitContent += message + "\n"

	commitHash := utility.Sha1([]byte(commitContent))
	commitPath := filepath.Join(".ritz", "objects", commitHash[:2])
	if err := os.MkdirAll(commitPath, 0755); err != nil {
		return "", err
	}
	commitFilePath := filepath.Join(commitPath, commitHash[2:])
	if err := utility.WriteObject(commitFilePath, []byte(commitContent)); err != nil {
		return "", err
	}
	return commitHash, nil
}

func updateRef(refName, commitHash string) error {
	refPath := filepath.Join(".ritz", refName)
	return ioutil.WriteFile(refPath, []byte(commitHash), 0644)
}

func getFilesFromIndex(index map[string]string) []string {
	var files []string
	for file := range index {
		files = append(files, file)
	}
	return files
}

