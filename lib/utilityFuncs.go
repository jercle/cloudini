package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

func UniqueNonEmptyElementsOf(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}
	return us
}

func SliceOfStringsToUnique(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, str := range slice {
		if _, ok := seen[str]; !ok {
			seen[str] = true
			result = append(result, str)
		}
	}
	return result
}

func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	// fmt.Println(typ)

	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		if strings.ToLower(fieldName) == "capabilities" {
			continue
		}
		fieldValueKind := val.Field(i).Kind()
		var fieldValue interface{}

		if fieldValueKind == reflect.Struct {
			fieldValue = StructToMap(val.Field(i).Interface())
		} else {
			fieldValue = val.Field(i).Interface()
		}

		result[fieldName] = fieldValue
	}

	return result
}

func AddJsonOmitemptyTagsToStructFile(path string, overwriteFile bool) (processedStruct string) {
	readFile, err := os.Open(path)
	CheckFatalError(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var processedLines []string

	for _, line := range fileLines {
		lineSplit := strings.SplitN(line, "`", 2)
		if len(lineSplit) > 1 && strings.Contains(line, "json:") && !strings.Contains(line, "omitempty") {
			tagString := strings.ReplaceAll(lineSplit[1], "`", "")
			// fmt.Println(tagString)
			tagString = tagString[:len(tagString)-1] + "\""
			jsonTag := tagString[5 : len(tagString)-1]
			// fmt.Println(jsonTag)
			jsonTag = "json:" + strings.ToLower(jsonTag[:1]) + jsonTag[1:] + ",omitempty,omitzero\"`"
			jsonTag = strings.ReplaceAll(jsonTag, ",attr", "")
			tagString = "`" + jsonTag
			processedLines = append(processedLines, lineSplit[0]+tagString)
		} else {
			processedLines = append(processedLines, line)
		}
	}
	processedFile := strings.Join(processedLines, "\n")
	formattedAndProcessed, err := format.Source([]byte(processedFile))
	CheckFatalError(err)

	if overwriteFile {
		err := os.WriteFile(path, formattedAndProcessed, 0644)
		CheckFatalError(err)
		return string(formattedAndProcessed)
	} else {
		return string(formattedAndProcessed)
	}
}

func AddBsonTagsFromJsonTagsToStructFile(path string, overwriteFile bool) (processedStruct string) {
	readFile, err := os.Open(path)
	CheckFatalError(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var processedLines []string

	for _, line := range fileLines {
		lineSplit := strings.SplitN(line, "`", 2)
		if len(lineSplit) > 1 && strings.Contains(line, "json:") && !strings.Contains(line, "bson:") {
			tagString := strings.ReplaceAll(lineSplit[1], "`", "")
			tagString = tagString[:len(tagString)-1] + "\""
			bsonTag := strings.ReplaceAll(tagString, "json:", "bson:")
			jsonTag := tagString[5:]
			tagString = "`json:" + jsonTag + " " + bsonTag + "`"
			processedLines = append(processedLines, lineSplit[0]+tagString)
		} else {
			processedLines = append(processedLines, line)
		}
	}
	processedFile := strings.Join(processedLines, "\n")
	formattedAndProcessed, err := format.Source([]byte(processedFile))
	CheckFatalError(err)

	if overwriteFile {
		err := os.WriteFile(path, formattedAndProcessed, 0644)
		CheckFatalError(err)
		return string(formattedAndProcessed)
	} else {
		return string(formattedAndProcessed)
	}
}

func ChangeJsonIdToMongoIdInFile(path string, overwriteFile bool) (processedStruct string) {
	readFile, err := os.Open(path)
	CheckFatalError(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var processedLines []string

	for _, line := range fileLines {
		// lineSplit := strings.SplitN(line, "`", 2)
		// if len(lineSplit) > 1 && strings.Contains(line, "json:") && !strings.Contains(line, "bson:") {
		// 	tagString := strings.ReplaceAll(lineSplit[1], "`", "")
		// 	tagString = tagString[:len(tagString)-1] + "\""
		// 	bsonTag := strings.ReplaceAll(tagString, "json:", "bson:")
		// 	jsonTag := tagString[5:]
		// 	tagString = "`json:" + jsonTag + " " + bsonTag + "`"
		// 	processedLines = append(processedLines, lineSplit[0]+tagString)
		// } else {
		// 	processedLines = append(processedLines, line)
		// }
		newLine := strings.ReplaceAll(line, "\"id\"", "\"_id\"")
		processedLines = append(processedLines, newLine)
	}
	processedFile := strings.Join(processedLines, "\n")
	// formattedAndProcessed, err := format.Source([]byte(processedFile))
	// CheckFatalError(err)

	if overwriteFile {
		err := os.WriteFile(path, []byte(processedFile), 0644)
		CheckFatalError(err)
		return string(processedFile)
	} else {
		return string(processedFile)
	}
}

func ChangeCaseOfJsonAndBsonTags(path string, overwriteFile bool) (processedStruct string) {
	readFile, err := os.Open(path)
	CheckFatalError(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var processedLines []string

	for _, line := range fileLines {
		if !strings.Contains(line, "json:\"") && !strings.Contains(line, "bson:\"") {
			processedLines = append(processedLines, line)
		} else {
			jsonSplit := strings.SplitN(line, "json:\"", 2)
			jsonTag := strings.ToLower(jsonSplit[1][:1]) + jsonSplit[1][1:]
			jsonSplit[1] = jsonTag
			jsonFixed := strings.Join(jsonSplit, "json:\"")
			bsonSplit := strings.SplitN(jsonFixed, "bson:\"", 2)
			bsonTag := strings.ToLower(bsonSplit[1][:1]) + bsonSplit[1][1:]
			bsonSplit[1] = bsonTag
			fixedLine := strings.Join(bsonSplit, "bson:\"")
			processedLines = append(processedLines, fixedLine)
		}
		// lineSplit := strings.SplitN(line, "`", 2)
		// if len(lineSplit) > 1 && !strings.Contains(line, "json:") || !strings.Contains(line, "bson:") {
		// 	tagString := strings.ReplaceAll(lineSplit[1], "`", "")
		// 	tagString = tagString[:len(tagString)-1] + "\""
		// 	bsonTag := strings.ReplaceAll(tagString, "json:", "bson:")
		// 	jsonTag := tagString[5:]
		// 	tagString = "`json:" + jsonTag + " " + bsonTag + "`"
		// 	processedLines = append(processedLines, lineSplit[0]+tagString)
		// } else {
		// 	processedLines = append(processedLines, line)
		// }
	}
	processedFile := strings.Join(processedLines, "\n")
	formattedAndProcessed, err := format.Source([]byte(processedFile))
	CheckFatalError(err)

	if overwriteFile {
		err := os.WriteFile(path, formattedAndProcessed, 0644)
		CheckFatalError(err)
		return string(formattedAndProcessed)
	} else {
		return string(formattedAndProcessed)
	}
}

func ProgressBar(count int, itsString string, barNum int, barTotal int, description string) *progressbar.ProgressBar {
	if itsString == "" {
		itsString = "it"
	}
	var barDesc string
	if barTotal > 1 {
		barDesc = "[cyan][" + strconv.Itoa(barNum) + "/" + strconv.Itoa(barTotal) + "][reset] " + description
	} else {
		barDesc = description
	}
	return progressbar.NewOptions(count,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetItsString(itsString),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionThrottle(100*time.Millisecond),
		// progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription(barDesc),
	)
}

func GetDataFromPackerLogfile(path string) *PackerLogBuildData {
	readFile, err := os.Open(path)
	CheckFatalError(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	var (
		buildCompleteLine int
		logData           PackerLogBuildData
		imgRef            *AzureResourceStorageProfileImageReference
		outImageIdFound   bool
	)
	for i, line := range fileLines {
		if strings.Contains(line, "Publishing to Shared Image Gallery") {
			buildCompleteLine = i
		}

		if strings.Contains(line, "performing CreateOrUpdate: unexpected status") {
			return nil
		}
	}
	if buildCompleteLine == 0 {
		return nil
	}

	for _, line := range fileLines {

		if logData.BuildBaseImageVersion == nil && strings.Contains(line, "exactVersion") {
			bodyStr := strings.Split(strings.Split(line, "body=\"")[1], "\" status=")[0]
			bodyStr = strings.ReplaceAll(bodyStr, "\\r", "")
			bodyStr = strings.ReplaceAll(bodyStr, "\\n", "")
			bodyStr = strings.ReplaceAll(bodyStr, "\\\\\"", "'")
			bodyStr = strings.ReplaceAll(bodyStr, "\\", "")

			if bodyStr[len(bodyStr)-1:] == "\"" {
				bodyStr = strings.TrimSuffix(bodyStr, "\"")
			}

			var bodyData AzureResourceDetails
			err = json.Unmarshal([]byte(bodyStr), &bodyData)
			CheckFatalError(err)

			// props := *bodyData.Properties

			imgRef = bodyData.Properties.StorageProfile.ImageReference

			// JsonMarshalAndPrint(props)
			// os.Exit(0)
			// imgRef = props["StorageProfile"]["ImageReference"]

			if strings.HasSuffix(imgRef.ID, "latest") {
				imgRef.ID = strings.TrimSuffix(imgRef.ID, "latest") + imgRef.ExactVersion
			}

			logData.BuildBaseImageVersion = imgRef
		}

		if strings.Contains(line, "SIG publish resource group     ") {
			logData.OutputImgResGrp, _ = GetStringInBetween(line, "'", "'")
		}

		if strings.HasPrefix(line, "ManagedImageSharedImageGalleryId") {
			outImgVersId := strings.Split(line, " ")[1]
			logData.OutputImgId = outImgVersId
			// logData.OutputImgId = strings.Split(outImgVersId, "/versions/")[0]
			outImageIdFound = true
		}

		if strings.Contains(line, "SIG gallery name     ") {
			logData.OutputImgGalleryName, _ = GetStringInBetween(line, "'", "'")
		}

		if strings.Contains(line, "SIG image name     ") {
			logData.OutputImgDef, _ = GetStringInBetween(line, "'", "'")
		}

		if strings.Contains(line, "SIG image version     ") {
			logData.OutputImgVersion, _ = GetStringInBetween(line, "'", "'")
		}

		if strings.Contains(line, "starting builder arm") {
			bistSplit := strings.Split(line, " ")[0:2]
			bist := strings.Join(bistSplit, " ")
			logData.BuildImageStartTime, err = time.Parse("2006/01/02 15:04:05", bist)
			CheckFatalError(err)
		}
	}

	if !outImageIdFound {
		for _, line := range fileLines {

			if strings.Contains(line, "Microsoft.Compute/galleries/images/versions") && strings.Contains(line, "method=\"PUT\"") {
				bodyStr := strings.Split(strings.Split(line, "body=\"")[1], "\" status=")[0]
				bodyStr = strings.ReplaceAll(bodyStr, "\\r", "")
				bodyStr = strings.ReplaceAll(bodyStr, "\\n", "")
				bodyStr = strings.ReplaceAll(bodyStr, "\\\\\"", "'")
				bodyStr = strings.ReplaceAll(bodyStr, "\\", "")

				if bodyStr[len(bodyStr)-1:] == "\"" {
					bodyStr = strings.TrimSuffix(bodyStr, "\"")
				}

				var bodyData PackerPublishImageResponse
				err = json.Unmarshal([]byte(bodyStr), &bodyData)
				CheckFatalError(err)

				logData.OutputImgId = bodyData.ID
				// logData.OutputImgId = strings.Split(bodyData.ID, "/versions/")[0]
			}
		}
	}

	bastSplit := strings.Split(fileLines[0], " ")[0:2]
	bast := strings.Join(bastSplit, " ")
	logData.AzDoStartTime, err = time.Parse("2006/01/02 15:04:05", bast)
	CheckFatalError(err)

	bactSplit := strings.Split(fileLines[len(fileLines)-1], " ")[0:2]
	bact := strings.Join(bactSplit, " ")
	logData.AzDoCompleteTime, err = time.Parse("2006/01/02 15:04:05", bact)
	CheckFatalError(err)

	bictSplit := strings.Split(fileLines[buildCompleteLine+1], " ")[0:2]
	bict := strings.Join(bictSplit, " ")
	logData.BuildImageCompleteTime, err = time.Parse("2006/01/02 15:04:05", bict)
	CheckFatalError(err)

	logData.AzDoDuration = logData.AzDoCompleteTime.Sub(logData.AzDoStartTime)
	logData.BuildImageDuration = logData.BuildImageCompleteTime.Sub(logData.BuildImageStartTime)

	dir, file := filepath.Split(path)
	bnSplit := strings.Split(file, "-")
	logData.AzDoBuildName = bnSplit[1]
	logData.BuildImageEnv = strings.Split(bnSplit[2], ".")[0]
	logData.AzDoHost = filepath.Base(dir)
	logData.AzDoLogFile = logData.AzDoHost + "/" + file

	return &logData
}

func GetDataFromMultiplePackerLogFiles(basePath string) (buildData []PackerLogBuildData) {
	paths := GetFullFilePaths(basePath)

	for _, path := range paths {
		logData := GetDataFromPackerLogfile(path)
		if logData != nil {
			buildData = append(buildData, *logData)
		}
	}

	return
}

func GetStringInBetween(str string, startS string, endS string) (result string, found bool) {
	s := strings.Index(str, startS)
	if s == -1 {
		return result, false
	}
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	if e == -1 {
		return result, false
	}
	result = newS[:e]
	return result, true
}

//
//

func AddXmlTagsFromJsonTagsToStructFile(path string, overwriteFile bool) (processedStruct string) {
	readFile, err := os.Open(path)
	CheckFatalError(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var processedLines []string

	for _, line := range fileLines {
		lineSplit := strings.SplitN(line, "`", 2)
		if len(lineSplit) > 1 && strings.Contains(line, "json:") && !strings.Contains(line, "xml:") {
			tagString := strings.ReplaceAll(lineSplit[1], "`", "")
			tagString = tagString[:len(tagString)-1] + "\""
			xmlTag := strings.ReplaceAll(tagString, "json:", "xml:")
			jsonTag := tagString[5:]
			tagString = "`json:" + jsonTag + " " + xmlTag + "`"
			processedLines = append(processedLines, lineSplit[0]+tagString)
		} else {
			processedLines = append(processedLines, line)
		}
	}
	processedFile := strings.Join(processedLines, "\n")
	formattedAndProcessed, err := format.Source([]byte(processedFile))
	CheckFatalError(err)

	if overwriteFile {
		err := os.WriteFile(path, formattedAndProcessed, 0644)
		CheckFatalError(err)
		return string(formattedAndProcessed)
	} else {
		return string(formattedAndProcessed)
	}
}

//
//

func PrintSliceStringsWithIndexes(slice []string) {
	for i, element := range slice {
		fmt.Println(i, element)
	}
}

func PrintSliceIntsWithIndexes(slice []int) {
	for i, element := range slice {
		fmt.Println(i, element)
	}
}

//
//

func IsSpecificDay(t time.Time, year int, month time.Month, day int) bool {
	return t.Year() == year && t.Month() == month && t.Day() == day
}

//
//

func SelectMapStringFieldsFromArrayOfKeys(dataMap map[string]string, keysToSelect []string) map[string]string {
	newMap := make(map[string]string)
	for _, key := range keysToSelect {
		// Check if the key exists in the map
		if _, ok := dataMap[key]; ok {
			// fmt.Printf("Key: %s, Value: %d\n", key, value)
			newMap[key] = dataMap[key]
		}
	}
	return newMap
}

//
//

func HttpBearerGet(urlString string, bearerToken string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	CheckFatalError(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+bearerToken)

	res, err := http.DefaultClient.Do(req)
	CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)

	if res.StatusCode == 404 {
		jsonErr := `{"status": "` + res.Status + `", "response": ` + string(responseBody) + `}`
		err = fmt.Errorf(jsonErr)
		return nil, err
	}

	return responseBody, nil
}

//
//

func LoggerJson() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
