package lib

import (
	"encoding/json"
	"fmt"
	"log"
)

// Simple wrapper to
//
//	if err!=nil {
//	  log.Fatal(err)
//	}
func CheckFatalError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Simple wraper around
//
// fmt.Println((string(jsonStr))
func PrintJsonBytes(jsonBytes []byte) {
	fmt.Println(string(jsonBytes))
}

func MarshalAndPrintJson(data any) {
	jsonStr, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonStr))
}
