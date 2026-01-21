package lib

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
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

//
//

func RemoveSubdirectoriesOfPath(path string) {
	dirs, err := os.ReadDir(path)
	CheckFatalError(err)
	for _, dir := range dirs {
		fullDir := filepath.Join(path, dir.Name())
		os.RemoveAll(fullDir)
	}
}

//
//

// func GetSubdirs(path string) []string {
// 	var (
// 		fullFilePaths []string
// 	)

// 	err := filepath.WalkDir(path,
// 		func(path string, info os.FileInfo, err error) error {
// 			CheckFatalError(err)
// 			// if !info.IsDir() {
// 			fullFilePaths = append(fullFilePaths, info.Name())
// 			// }
// 			return nil
// 		})
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return fullFilePaths
// }

//
//

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

func RemoveJsonByteOrderMark(str []byte) []byte {
	processed := bytes.TrimPrefix(str, []byte("\xef\xbb\xbf"))
	return processed
}

func SplitPath(path string) []string {
	subPath := path
	var result []string
	for {
		subPath = filepath.Clean(subPath) // Amongst others, removes trailing slashes (except for the root directory).

		dir, last := filepath.Split(subPath)
		if last == "" {
			if dir != "" { // Root directory.
				result = append(result, dir)
			}
			break
		}
		result = append(result, last)

		if dir == "" { // Nothing to split anymore.
			break
		}
		subPath = dir
	}

	slices.Reverse(result)
	return result
}

func ReadFileUTF16(filename string) ([]byte, error) {

	// Read the file into a []byte:
	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Make an tranformer that converts MS-Win default to UTF8:
	win16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	// Make a transformer that is like win16be, but abides by BOM:
	utf16bom := unicode.BOMOverride(win16be.NewDecoder())

	// Make a Reader that uses utf16bom:
	unicodeReader := transform.NewReader(bytes.NewReader(raw), utf16bom)

	// decode and print:
	decoded, err := io.ReadAll(unicodeReader)
	return decoded, err
}
