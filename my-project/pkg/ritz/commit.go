package ritz

import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "time"

    "/my-project/pkg/utility"
)

func main() {
    // Parse command-line flags
    commitMessage := flag.String("m", "Initial commit", "Commit message")
    flag.Parse()

    // Check if the repository is initialized
    if !isInitialized() {
        fmt.Println("The directory is not a ritz directory. Please initialize it first.")
        return
    }

    // Prepare the tree content
    treeContent, err := prepareTreeContent()
    if err != nil {
        fmt.Println("Failed to prepare tree content:", err)
        return
    }

    // Create the commit content
    commitContent := createCommitContent(treeContent, *commitMessage)

    // Create the commit
    if err := createCommit(commitContent); err != nil {
        fmt.Println("Failed to create commit:", err)
    }
}

func isInitialized() bool {
    _, err := os.Stat(".ritz")
    return err == nil
}

func prepareTreeContent() (string, error) {
    // List files in the current directory
    files, err := utility.ListFiles(".")
    if err != nil {
        return "", err
    }

    var treeContent string
    for _, file := range files {
        // Exclude directories and commit.go file itself
        if !file.IsDir() && file.Name() != "commit.go" {
            // Read file content
            fileContent, err := ioutil.ReadFile(file.Name())
            if err != nil {
                return "", err
            }

            // Compute SHA-1 hash of file content
            hash := utility.Sha1Hash(fileContent)

            // Store file content as a blob
            blobPath := filepath.Join(".ritz", "objects", hash)
            compressedContent, err := utility.Compress(fileContent)
            if err != nil {
                return "", err
            }
            if err := ioutil.WriteFile(blobPath, compressedContent, 0644); err != nil {
                return "", err
            }

            // Append to tree content
            treeContent += fmt.Sprintf("%s %s\n", hash, file.Name())
        }
    }
    return treeContent, nil
}

func createCommitContent(treeContent, message string) string {
    // Get current time in Unix format
    commitTime := time.Now().Unix()

    // Compute SHA-1 hash of tree content
    treeHash := utility.Sha1Hash([]byte(treeContent))

    // Define author information
    author := "Initial Commit <anonymous@example.com>"

    // Construct commit content
    return fmt.Sprintf("tree %s\nauthor: %s %d +0000\ncommitter: %s %d +0000\n\n%s\n",
        treeHash, author, commitTime, author, commitTime, message)
}

func createCommit(commitContent string) error {
    // Compute SHA-1 hash of commit content
    commitHash := utility.Sha1Hash([]byte(commitContent))

    // Write commit content to objects directory
    commitPath := filepath.Join(".ritz", "objects", commitHash)
    if err := ioutil.WriteFile(commitPath, []byte(commitContent), 0644); err != nil {
        return err
    }

    // Update HEAD to point to new commit
    headContent := fmt.Sprintf("ref: refs/heads/master\n%s\n", commitHash)
    if err := ioutil.WriteFile(".ritz/HEAD", []byte(headContent), 0644); err != nil {
        return err
    }

    return nil
}

