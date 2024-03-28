package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
)

// Simple wrapper to
//
//	if err!=nil {
//	  log.Fatal(err)
//	}
func CheckFatalError(err error) {
	if err != nil {
		_, file, no, _ := runtime.Caller(1)
		log.Fatalln(file+":"+strconv.Itoa(no)+":0", err)
	}
}

func PrintSrcLoc(str ...string) {
	_, file, no, _ := runtime.Caller(1)
	fmt.Println(file+":"+strconv.Itoa(no)+":0", str)
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

type WaitGroupCount struct {
	sync.WaitGroup
	count int64
}

func (wg *WaitGroupCount) Add(delta int) {
	atomic.AddInt64(&wg.count, int64(delta))
	wg.WaitGroup.Add(delta)
}

func (wg *WaitGroupCount) Done() {
	atomic.AddInt64(&wg.count, -1)
	wg.WaitGroup.Done()
}

func (wg *WaitGroupCount) GetCount() int {
	return int(atomic.LoadInt64(&wg.count))
}
