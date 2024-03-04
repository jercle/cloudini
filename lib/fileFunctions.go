package lib

import (
	"log"
	"os"
	"path/filepath"
)

func GetFullFilePaths(path string) []string {
	var (
		fullFilePaths []string
	)

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				fullFilePaths = append(fullFilePaths, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return fullFilePaths
}
