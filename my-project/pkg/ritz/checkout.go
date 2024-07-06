package ritz

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

func CheckoutBranch(branchName string) error {
    branchPath := filepath.Join(branchDir, branchName)
    if _, err := os.Stat(branchPath); os.IsNotExist(err) {
        return fmt.Errorf("branch '%s' does not exist", branchName)
    }

    // Read the current branch from HEAD
    headContent, err := ioutil.ReadFile(".ritz/HEAD")
    if err != nil {
        return err
    }
    currentBranch := strings.TrimPrefix(strings.TrimSpace(string(headContent)), "ref: refs/heads/")

    if branchName == currentBranch {
        fmt.Printf("Already on '%s'\n", branchName)
        return nil
    }

    // Update HEAD to point to the new branch
    newHeadContent := fmt.Sprintf("ref: refs/heads/%s", branchName)
    err = ioutil.WriteFile(".ritz/HEAD", []byte(newHeadContent), 0644)
    if err != nil {
        return err
    }

    return nil
}

func CreateAndCheckoutBranch(branchName string) error {
    err := CreateBranch(branchName)
    if err != nil {
        return err
    }

    return CheckoutBranch(branchName)
}

