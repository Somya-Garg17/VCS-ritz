package ritz

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"my-project/pkg/utility"
)

func Commit(message string) {
	if !utility.IsInitialized() {
		fmt.Println("This is not a ritz repository. Run 'ritz init' to initialize one.")
		return
	}

	indexPath := filepath.Join(".ritz", "index")
	index := utility.LoadIndex(indexPath)

	if len(index) == 0 {
		untrackedFiles := getUntrackedFiles()
		fmt.Println("On branch master")
		fmt.Println("\nInitial commit\n")
		if len(untrackedFiles) > 0 {
			fmt.Println("Untracked files:")
			fmt.Println("  (use \"ritz add <file>...\" to include in what will be committed)")
			for _, file := range untrackedFiles {
				fmt.Println("\t" + file)
			}
			fmt.Println("\nnothing added to commit but untracked files present (use \"ritz add\" to track)")
		} else {
			fmt.Println("nothing to commit, working tree clean")
		}
		return
	}

	tree := createTreeObject(index)
	treeHash := utility.Sha1([]byte(tree))
	treePath := filepath.Join(".ritz", "objects", treeHash[:2], treeHash[2:])
	if err := utility.WriteObject(treePath, []byte(tree)); err != nil {
		fmt.Println("Error writing tree object:", treePath, err)
		return
	}

	headPath := filepath.Join(".ritz", "HEAD")
	headRef := readHead(headPath)

	commitContent := fmt.Sprintf(
		"tree %s\nparent %s\nauthor Your Name <you@example.com> %d +0000\ncommitter Your Name <you@example.com> %d +0000\n\n%s\n",
		treeHash, headRef, time.Now().Unix(), time.Now().Unix(), message,
	)
	commitHash := utility.Sha1([]byte(commitContent))
	commitPath := filepath.Join(".ritz", "objects", commitHash[:2], commitHash[2:])
	if err := utility.WriteObject(commitPath, []byte(commitContent)); err != nil {
		fmt.Println("Error writing commit object:", commitPath, err)
		return
	}

	if err := updateRef(headPath, "refs/heads/master", commitHash); err != nil {
		fmt.Println("Error updating HEAD:", err)
		return
	}

	fmt.Printf("[master (root-commit) %s] %s\n", commitHash, message)

	// Display changes summary
	fmt.Printf(" %d file(s) changed\n", len(index))
	for file := range index {
		fmt.Printf(" create mode 100644 %s\n", file)
	}
}

func createTreeObject(index map[string]string) string {
	var lines []string
	for file, hash := range index {
		lines = append(lines, fmt.Sprintf("100644 blob %s\t%s", hash, file))
	}
	return strings.Join(lines, "\n")
}

func readHead(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(content))
}

func updateRef(headPath, refPath, commitHash string) error {
	refPath = filepath.Join(".ritz", refPath)
	if err := os.WriteFile(refPath, []byte(commitHash), 0644); err != nil {
		return err
	}
	return os.WriteFile(headPath, []byte("ref: "+refPath), 0644)
}

func getUntrackedFiles() []string {
	var untrackedFiles []string
	files, _ := ioutil.ReadDir(".")
	for _, file := range files {
		if file.IsDir() || file.Name() == ".ritz" {
			continue
		}
		untrackedFiles = append(untrackedFiles, file.Name())
	}
	return untrackedFiles
}

