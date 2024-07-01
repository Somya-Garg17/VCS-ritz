package utility

import (
        "compress/zlib"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadObject(hash string) ([]byte, error) {
	objectPath := filepath.Join(".ritz", "objects", hash[:2], hash[2:])
	file, err := os.Open(objectPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	zr, err := zlib.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	content, err := ioutil.ReadAll(zr)
	if err != nil {
		return nil, err
	}
	return content, nil
}

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

func GetUntrackedFiles(index map[string]string) ([]string, error) {
	var untrackedFiles []string

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasPrefix(path, ".ritz") {
			return nil
		}
		if _, tracked := index[path]; !tracked {
			untrackedFiles = append(untrackedFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return untrackedFiles, nil
}
