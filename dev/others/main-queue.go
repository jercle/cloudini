// https://www.opsdash.com/blog/job-queues-in-go.html
// https://medium.com/@basakabhijoy/job-queue-in-go-20a378a6c7c1
// https://gist.github.com/harlow/dbcd639cf8d396a2ab73
// https://github.com/hibiken/asynq
// https://reintech.io/blog/implementing-distributed-task-queue-go#google_vignette

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func process(jobChannel chan int, worker int) {
	defer wg.Done()
	for job := range jobChannel {
		fmt.Printf("Job %v picked up by %v\n", job, worker)
	}
}

func main() {
	jobChannel := make(chan int, 10)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go process(jobChannel, i)
	}

	for i := 0; i < 1000; i++ {
		jobChannel <- i
	}

	close(jobChannel)

	wg.Wait()
}
