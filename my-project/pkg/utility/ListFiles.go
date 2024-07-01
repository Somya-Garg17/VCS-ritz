package utility

import (
    "io/ioutil"
    "os"
)

func ListFiles(dir string) ([]os.FileInfo, error) {
    return ioutil.ReadDir(dir)
}

