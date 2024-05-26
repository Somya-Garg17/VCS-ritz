package ritz

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"my-project/pkg/utility"
)

func Commit() {
	commitMessage := flag.String("m", "Initial commit", "Commit message")
	flag.Parse()

	if !utility.IsInitialized() {
		fmt.Println("The directory is not initialized as a ritz directory. Please run 'ritz init' first.")
		return
	}

	treeContent, err := prepareTreeContent()
	if err != nil {
		fmt.Println("Failed to prepare tree content:", err)
		return
	}

	treeHash := utility.Sha1([]byte(treeContent))
	treePath := filepath.Join(".ritz", "objects", treeHash)
	if err := ioutil.WriteFile(treePath, []byte(treeContent), 0644); err != nil {
		fmt.Println("Failed to write tree object:", err)
		return
	}

	commitContent := createCommitContent(treeHash, *commitMessage)
	if err := createCommit(commitContent, *commitMessage); err != nil {
		fmt.Println("Failed to create commit:", err)
	}

	if err := VerifyCommit(); err != nil {
		fmt.Println("Failed to verify commit:", err)
	}
}

func prepareTreeContent() (string, error) {
	files, err := utility.ListFiles(".")
	if err != nil {
		return "", err
	}

	var treeContent strings.Builder
	for _, file := range files {
		if !file.IsDir() && file.Name() != "main.go" && file.Name() != "my_cli" && file.Name() != "go.mod" && file.Name() != "Inflate" && file.Name() !="Inflate.go" {
			fileContent, err := ioutil.ReadFile(file.Name())
			if err != nil {
				return "", err
			}

			hash := utility.Sha1(fileContent)

			blobPath := filepath.Join(".ritz", "objects", hash)
			compressedContent, err := utility.Compress(fileContent)
			if err != nil {
				return "", err
			}
			if err := ioutil.WriteFile(blobPath, compressedContent, 0644); err != nil {
				return "", err
			}

			fmt.Fprintf(&treeContent, "%s %s\n", hash, file.Name())
		}
	}
	return treeContent.String(), nil
}

func createCommitContent(treeHash, message string) string {
	commitTime := time.Now().Unix()
	author := "Initial Commit <anonymous@example.com>"
	return fmt.Sprintf("tree %s\nauthor: %s %d +0000\ncommitter: %s %d +0000\n\n%s\n",
		treeHash, author, commitTime, author, commitTime, message)
}

func createCommit(commitContent, commitMessage string) error {
	commitHash := utility.Sha1([]byte(commitContent))

	commitPath := filepath.Join(".ritz", "objects", commitHash)
	if err := ioutil.WriteFile(commitPath, []byte(commitContent), 0644); err != nil {
		return err
	}

	if err := ioutil.WriteFile(".ritz/refs/heads/master", []byte(commitHash), 0644); err != nil {
		return err
	}

	return nil
}

func VerifyCommit() error {
	headContent, err := ioutil.ReadFile(".ritz/refs/heads/master")
	if err != nil {
		return fmt.Errorf("failed to read HEAD reference: %v", err)
	}
	commitHash := strings.TrimSpace(string(headContent))

	commitPath := filepath.Join(".ritz", "objects", commitHash)
	commitContent, err := ioutil.ReadFile(commitPath)
	if err != nil {
		return fmt.Errorf("failed to read commit object: %v", err)
	}

	commitHashComputed := utility.Sha1(commitContent)

	if commitHash != commitHashComputed {
		return fmt.Errorf("commit verification failed: hash mismatch")
	}

	fmt.Println("Commit verification successful.")
	return nil
}

