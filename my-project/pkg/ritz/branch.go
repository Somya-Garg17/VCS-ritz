package ritz

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

const branchDir = ".ritz/refs/heads/"

// ListBranches lists all the branches
func ListBranches() error {
    files, err := ioutil.ReadDir(branchDir)
    if err != nil {
        return err
    }

    // Read the current branch from HEAD
    headContent, err := ioutil.ReadFile(".ritz/HEAD")
    if err != nil {
        return err
    }
    currentBranch := strings.TrimPrefix(strings.TrimSpace(string(headContent)), "ref: refs/heads/")

    for _, file := range files {
        branchName := file.Name()
        if branchName == currentBranch {
            fmt.Printf("* %s\n", branchName)
        } else {
            fmt.Println(branchName)
        }
    }
    return nil
}

// CreateBranch creates a new branch
func CreateBranch(branchName string) error {
    branchPath := filepath.Join(branchDir, branchName)
    if _, err := os.Stat(branchPath); err == nil {
        return fmt.Errorf("branch '%s' already exists", branchName)
    }

    currentHead, err := ioutil.ReadFile(".ritz/HEAD")
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(branchPath, currentHead, 0644)
    if err != nil {
        return err
    }

    return nil
}

// DeleteBranch deletes an existing branch
func DeleteBranch(branchName string) error {
    branchPath := filepath.Join(branchDir, branchName)
    if _, err := os.Stat(branchPath); os.IsNotExist(err) {
        return fmt.Errorf("branch '%s' does not exist", branchName)
    }

    return os.Remove(branchPath)
}

// RenameBranch renames an existing branch
func RenameBranch(oldName, newName string) error {
    oldPath := filepath.Join(branchDir, oldName)
    newPath := filepath.Join(branchDir, newName)

    if _, err := os.Stat(oldPath); os.IsNotExist(err) {
        return fmt.Errorf("branch '%s' does not exist", oldName)
    }
    if _, err := os.Stat(newPath); err == nil {
        return fmt.Errorf("branch '%s' already exists", newName)
    }

    return os.Rename(oldPath, newPath)
}

