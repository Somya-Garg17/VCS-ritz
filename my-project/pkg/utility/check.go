package utility

import (
	"os"
)

func IsInitialized() bool {
	_, err := os.Stat(".ritz")
	return err == nil
}

