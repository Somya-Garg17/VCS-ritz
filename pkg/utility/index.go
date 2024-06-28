package utility

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func LoadIndex(path string) map[string]string {
	index := make(map[string]string)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return index
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		index[parts[1]] = parts[0]
	}
	return index
}

func SaveIndex(path string, index map[string]string) error {
	var lines []string
	for file, hash := range index {
		lines = append(lines, fmt.Sprintf("%s %s", hash, file))
	}
	return ioutil.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}

func WriteObject(path string, content []byte) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	compressedContent, err := Compress(content)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, compressedContent, 0644)
}

