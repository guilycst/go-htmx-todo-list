package dirutil

import (
	"fmt"
	"os"
)

func IsDir(dir string) error {
	file, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("%s is not a dir", dir)
	}

	return nil
}
