package system

import (
	"os"
	"path/filepath"
)

func GetSymfonyProjectDir() (string, error) {
	execDir, err := GetExecDir()
	path := os.Args[1]
	if !filepath.IsAbs(path) {
		path = filepath.Clean(execDir + "/" + path)
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return "", err
	}

	return path, err
}

func GetExecDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path, nil
}
