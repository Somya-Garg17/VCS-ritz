package ritz

import (
	"fmt"
	"path/filepath"
	"strings"
	"io/ioutil"

	"my-project/pkg/utility"
)

func Status() {
	if !utility.IsInitialized() {
		fmt.Println("This is not a ritz repository. Run 'ritz init' to initialize one.")
		return
	}

	indexPath := filepath.Join(".ritz", "index")
	index := utility.LoadIndex(indexPath)

	headCommit := getHeadCommit()
	if headCommit == "" {
		fmt.Println("On branch master")
		fmt.Println("\nInitial commit\n")
	} else {
		fmt.Println("On branch master")
	}

	// Check for changes to be committed
	changesToBeCommitted := getChangesToBeCommitted(index, headCommit)
	if len(changesToBeCommitted) > 0 {
		fmt.Println("\nChanges to be committed:")
		fmt.Println("  (use \"ritz reset HEAD <file>...\" to unstage)")
		for _, file := range changesToBeCommitted {
			fmt.Println("\tnew file:", file)
		}
	} else {
		if headCommit == "" {
			untrackedFiles, err := utility.GetUntrackedFiles(index)
			if err != nil {
				fmt.Println("Error getting untracked files:", err)
				return
			}
			if len(untrackedFiles) > 0 {
				fmt.Println("\nUntracked files:")
				fmt.Println("  (use \"ritz add <file>...\" to include in what will be committed)")
				for _, file := range untrackedFiles {
					fmt.Println("\t" + file)
				}
			} else {
				fmt.Println("\nnothing added to commit but untracked files present (use \"ritz add\" to track)")
			}
		} else {
			fmt.Println("\nnothing to commit, working tree clean")
		}
	}

	// Check for changes not staged for commit
	changesNotStaged := getChangesNotStaged(index)
	if len(changesNotStaged) > 0 {
		fmt.Println("\nChanges not staged for commit:")
		fmt.Println("  (use \"ritz add <file>...\" to update what will be committed)")
		for _, file := range changesNotStaged {
			fmt.Println("\tmodified:", file)
		}
	}
}

func getHeadCommit() string {
	headPath := filepath.Join(".ritz", "HEAD")
	content, err := ioutil.ReadFile(headPath)
	if err != nil {
		return ""
	}
	refPath := strings.TrimSpace(strings.TrimPrefix(string(content), "ref: "))
	refPath = filepath.Join(".ritz", refPath)
	commitHash, err := ioutil.ReadFile(refPath)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(commitHash))
}

func getChangesToBeCommitted(index map[string]string, headCommit string) []string {
	var changes []string

	if headCommit == "" {
		for file := range index {
			changes = append(changes, file)
		}
		return changes
	}

	for file, indexHash := range index {
		headHash := getObjectHashFromCommit(headCommit, file)
		if headHash != indexHash {
			changes = append(changes, file)
		}
	}

	return changes
}

func getChangesNotStaged(index map[string]string) []string {
	var changes []string

	for file := range index {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}

		workDirHash := utility.Sha1(content)
		indexHash := index[file]
		if workDirHash != indexHash {
			changes = append(changes, file)
		}
	}

	return changes
}

func getObjectHashFromCommit(commitHash, filePath string) string {
	commitPath := filepath.Join(".ritz", "objects", commitHash[:2], commitHash[2:])
	commitContent, err := ioutil.ReadFile(commitPath)
	if err != nil {
		return ""
	}

	lines := strings.Split(string(commitContent), "\n")
	var treeHash string
	for _, line := range lines {
		if strings.HasPrefix(line, "tree ") {
			treeHash = strings.TrimPrefix(line, "tree ")
			break
		}
	}

	if treeHash == "" {
		return ""
	}

	return getObjectHashFromTree(treeHash, filePath)
}

func getObjectHashFromTree(treeHash, filePath string) string {
	treePath := filepath.Join(".ritz", "objects", treeHash[:2], treeHash[2:])
	treeContent, err := ioutil.ReadFile(treePath)
	if err != nil {
		return ""
	}

	lines := strings.Split(string(treeContent), "\n")
	for _, line := range lines {
		if strings.HasSuffix(line, "\t"+filePath) {
			parts := strings.Split(line, " ")
			if len(parts) > 2 {
				return parts[1]
			}
		}
	}

	return ""
}

