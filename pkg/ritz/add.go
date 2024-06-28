package ritz

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"my-project/pkg/utility"
)

func Add(files []string) {
	if !utility.IsInitialized() {
		fmt.Println("This is not a ritz repository. Run 'ritz init' to initialize one.")
		return
	}

	indexPath := filepath.Join(".ritz", "index")
	index := utility.LoadIndex(indexPath)

	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", file, err)
			continue
		}

		sha1Hash := utility.Sha1(content)
		objectPath := filepath.Join(".ritz", "objects", sha1Hash[:2], sha1Hash[2:])

		if err := utility.WriteObject(objectPath, content); err != nil {
			fmt.Println("Error writing object:", objectPath, err)
			continue
		}

		index[file] = sha1Hash
	}

	if err := utility.SaveIndex(indexPath, index); err != nil {
		fmt.Println("Error saving index:", err)
	}
}

