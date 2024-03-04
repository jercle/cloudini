package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-faker/faker/v4"
)

type Records struct {
	Records []struct {
		Category      `faker:"oneof: compute, network, storage"`
		MacAddress    string `faker:"mac_address"`
		OperationName string `faker:"username"`
		Properties    struct {
			Version uint `faker:"oneof: 15, 16"`
			Flows   []struct {
				Flows []struct {
					FlowTuples []string `faker:"ipv4"`
					Mac        string   `faker:"mac_address"`
				} `json:"flows"`
				Rule string `faker:"username"`
			} `json:"flows"`
		} `json:"properties"`
		ResourceID string `faker:"uuid_hyphenated"`
		SystemID   string `faker:"uuid_hyphenated"`
		Time       string `faker:"time"`
	} `json:"records"`
}

func Example_withTags() {

	a := Records{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)

}

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		fmt.Println("worker", id, "started  job", j)
// 		time.Sleep(time.Second)
// 		fmt.Println("worker", id, "finished job", j)
// 		results <- j * 2
// 	}
// }

// func main() {

// 	const numJobs = 500
// 	jobs := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	for w := 1; w <= 100; w++ {
// 		go worker(w, jobs, results)
// 	}

// 	for j := 1; j <= numJobs; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)

// 	for a := 1; a <= numJobs; a++ {
// 		<-results
// 	}
// }

func main() {
	// files, err := os.ReadDir("./nsglogs")
	var (
		fileNames []string
		// fileData  []Records
		// wg        sync.WaitGroup
		// fileData  []Records
		// wg        sync.WaitGroup
	)

	err := filepath.Walk("./nsglogs",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				fmt.Println(path, info.Size())
				fileNames = append(fileNames, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	for _, fileName := range fileNames {
		fmt.Println(fileName)
	}

	// func GetActiveSub() (string, error) {
	// 	subs, _ := getSubs()

	// 	for _, sub := range subs.Subscriptions {
	// 		if sub.IsDefault {
	// 			return sub.ID, nil
	// 		}
	// 	}

	// 	return "", fmt.Errorf("No default subscription")
	// }

	// for _, file := range fileNames {

	// 	wg.Add(1)

	// 	// fmt.Println(file)

	// 	os.ReadFile(file)

	// 	fileData = append(fileData)
	// }

	// wg.Wait()

}

// func readAndParseCsv(path string) ([][]string, error) {
// 	csvFile, err := os.Open(path)
// 	if err != nil {
// 		return nil, fmt.Errorf("error opening %s\n", path)
// 	}

// 	var rows [][]string

// 	reader := csv.NewReader(csvFile)
// 	reader.Comma = '\t'
// 	for {
// 		row, err := reader.Read()
// 		if err == io.EOF {
// 			break
// 		}

// 		if err != nil {
// 			return rows, fmt.Errorf("failed to parse csv: %s", err)
// 		}

// 		rows = append(rows, row)
// 	}

// 	return rows, nil
// }
