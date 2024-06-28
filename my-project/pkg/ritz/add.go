package ritz

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"my-project/pkg/utility"
)

func Add(paths []string) {
	if !utility.IsInitialized() {
		fmt.Println("This is not a ritz repository. Run 'ritz init' to initialize one.")
		return
	}

	indexPath := filepath.Join(".ritz", "index")
	index := utility.LoadIndex(indexPath)

	for _, path := range paths {
		err := addPathToIndex(path, index)
		if err != nil {
			fmt.Println("Error adding path:", path, err)
			continue
		}
	}

	if err := utility.SaveIndex(indexPath, index); err != nil {
		fmt.Println("Error saving index:", err)
	}
}

func addPathToIndex(path string, index map[string]string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return filepath.Walk(path, func(walkPath string, walkInfo os.FileInfo, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if walkInfo.IsDir() {
				return nil
			}
			return addFileToIndex(walkPath, index)
		})
	}

	return addFileToIndex(path, index)
}

func addFileToIndex(file string, index map[string]string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	sha1Hash := utility.Sha1(content)
	objectPath := filepath.Join(".ritz", "objects", sha1Hash[:2], sha1Hash[2:])

	if err := utility.WriteObject(objectPath, content); err != nil {
		return err
	}

	relativePath, err := filepath.Rel(".", file)
	if err != nil {
		return err
	}

	index[relativePath] = sha1Hash
	return nil
}

