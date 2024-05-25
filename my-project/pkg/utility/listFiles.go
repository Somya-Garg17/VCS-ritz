package utility

import "os"

func Ls() ([]string, error) {
  // returns working dir in string
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
  // Opens dir and return *file type
  // file is a struct defined in os pkg and various methods defined
	dir, _ := os.Open(wd)
  // ReadDir and return []fs.DirEntry
  // fs.DirEntry is an interface
  // To implement fs.DirEntry you need to implement 4 methods
  // Name() String, IsDir() bool, Type() FileMode, Info (FileInfo, error)
	files, _ := dir.ReadDir(0)
	list := []string{}  // slice of names of file
	for _, file := range files {
		if !file.IsDir() {
			list = append(list, file.Name())
		}
	}
	return list, nil
}
