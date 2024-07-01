package ritz

import (
    "fmt"
    "io/ioutil"
    "log"
    "path/filepath"

    "github.com/sergi/go-diff/diffmatchpatch"
    "my-project/pkg/utility"
)

const (
    ritzDir = ".ritz"
    objDir  = ritzDir + "/objects"
)

// Diff function compares the working directory file with the committed file
func Diff(filePath string) {
    // Check if the repository is initialized
    if !utility.IsInitialized() {
        log.Fatal("Ritz repository not initialized")
    }

    // Read the current version of the file
    currentContent, err := ioutil.ReadFile(filePath)
    if err != nil {
        log.Fatalf("Error reading current file: %v", err)
    }
    fmt.Println("Current file content:", string(currentContent))

    // Read the last committed version of the file
    committedContent, err := readCommittedFile(filePath)
    if err != nil {
        fmt.Println("Error reading committed file:", err)
        committedContent = []byte("(no committed version)")
    }
    fmt.Println("Committed file content:", string(committedContent))

    // Generate and display the diff
    generateDiff(string(committedContent), string(currentContent))
}

func readCommittedFile(filePath string) ([]byte, error) {
    indexPath := filepath.Join(".ritz", "index")
    index := utility.LoadIndex(indexPath)

    sha1, found := index[filePath]
    if !found {
        return nil, fmt.Errorf("file not found in index")
    }

    committedFilePath := filepath.Join(objDir, sha1[:2], sha1[2:])

    compressedData, err := ioutil.ReadFile(committedFilePath)
    if err != nil {
        return nil, err
    }

    decompressedData, err := utility.Decompress(compressedData)
    if err != nil {
        return nil, err
    }

    return decompressedData, nil
}

func generateDiff(oldText, newText string) {
    dmp := diffmatchpatch.New()
    diffs := dmp.DiffMain(oldText, newText, false)
    fmt.Println(dmp.DiffPrettyText(diffs))
}

