package lib

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/briandowns/spinner"
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

func CheckFatalErrorWithSpinner(err error, s *spinner.Spinner) {
	if err != nil {
		_, file, no, _ := runtime.Caller(1)
		s.Stop()
		log.Fatalln(file+":"+strconv.Itoa(no)+":0", err)
	}
}

// func CheckFatalError(err error) {
// 	if err != nil {
// 		pc := make([]uintptr, 10)
// 		n := runtime.Callers(0, pc)
// 		if n == 0 {
// 			// No PCs available. This can happen if the first argument to
// 			// runtime.Callers is large.
// 			//
// 			// Return now to avoid processing the zero Frame that would
// 			// otherwise be returned by frames.Next below.
// 			// os.Exit(0)
// 		}

// 		pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
// 		frames := runtime.CallersFrames(pc)

// 		for {
// 			frame, more := frames.Next()

// 			// Process this frame.
// 			//
// 			// To keep this example's output stable
// 			// even if there are changes in the testing package,
// 			// stop unwinding when we leave package runtime.
// 			if !strings.Contains(frame.File, "runtime/") {
// 				break
// 			}
// 			fmt.Printf("- more:%v | %s\n", more, frame.Function)
// 			// file, line := frame.Func.FileLine(frame.PC)
// 			// fmt.Println(file, line)

// 			// Check whether there are more frames to process after this one.
// 			if !more {
// 				break
// 			}
// 		}
// 		//  .Caller(1)
// 		// log.Fatalln(file+":"+strconv.Itoa(no)+":0", err)
// 		// jsonStr, _ := json.MarshalIndent(frames, "", "  ")
// 		// fmt.Println(string(jsonStr))
// 		fmt.Println(frames)
// 		// log.Fatalln(callers, err)
// 		// os.Exit(1)
// 	}
// }

//
//

func CheckHttpGetError(err error) {
	if err != nil {
		jsonStr := err.Error()
		var httpGetErr HttpGetError
		json.Unmarshal([]byte(jsonStr), &httpGetErr)
		httpGetErrStr, _ := json.Marshal(httpGetErr)
		if httpGetErr.Status == "404 Not Found" {
			fmt.Println(string(httpGetErrStr))
		} else {
			_, file, no, _ := runtime.Caller(1)
			log.Fatalln(file+":"+strconv.Itoa(no)+":0", err)
		}
	}
}

//
//

type HttpGetError struct {
	Response interface{} `json:"response"`
	Status   string      `json:"status"`
}

//
//

func PrintSrcLoc(str ...string) {
	_, file, no, _ := runtime.Caller(1)
	fmt.Println(file+":"+strconv.Itoa(no)+":0", strings.Join(str, " "))
}

// Simple wraper around
//
// fmt.Println((string(jsonStr))
func PrintJsonBytes(jsonBytes []byte) {
	fmt.Println(string(jsonBytes))
}

//
//

func MarshalAndPrintJson(data any) {
	jsonStr, err := json.Marshal(data, jsontext.WithIndent("  "))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonStr))
}

//
//

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

//
//

func JsonMarshalAndPrint(str interface{}) {
	jsonStr, _ := json.Marshal(str, jsontext.WithIndent("  "))
	fmt.Println(string(jsonStr))
}

//
//

func JsonMarshalAndWriteFile(str interface{}, outputFile string) {
	jsonStr, _ := json.Marshal(str, jsontext.WithIndent("  "))
	os.WriteFile(outputFile, jsonStr, 0644)
}
