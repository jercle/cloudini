package lib

import (
	"fmt"
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
			CheckFatalError(err)
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

func CheckDirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func DeleteFilesInDirMatchingString(dir string, strMatch string) {
	if !CheckDirExists(dir) {
		fmt.Println("Directory does not exist")
		os.Exit(1)
	}

	files, err := filepath.Glob(dir + strMatch)
	CheckFatalError(err)

	for _, f := range files {
		err := os.Remove(f)
		CheckFatalError(err)
	}
}
