package ff

import (
	"log"
	"os"
	"path/filepath"
)

func CreatePath(path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if mkdirErr := os.MkdirAll(dir, 0770); mkdirErr != nil {
			return mkdirErr
		}
	}
	return nil
}

func FileAppend(filePath, data string) {
	dir := filepath.Dir(filePath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		mkdirErr := os.MkdirAll(dir, 0770)
		if mkdirErr != nil {
			log.Printf("mkdir err: %s", mkdirErr)
		}
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("cmd: add error, %s", err)
	}
	defer f.Close()
	if _, err := f.WriteString(data); err != nil {
		log.Printf("append error: %s", err)
	}
}

func ReadString(path string) (string, error) {
	if absPath, err := filepath.Abs(path); err != nil {
		return "", err
	} else if bytes, err := os.ReadFile(absPath); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
