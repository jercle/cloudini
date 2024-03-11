package main

import "fmt"

func main() {

	channel := make(chan string, 100)
	go func() {
		channel <- "Hello World"
	}()
	go func() {
		channel <- "Hello World2"
	}()
	go func() {
		channel <- "Hello World3"
	}()
	go func() {
		channel <- "Hello World4"
	}()
	go func() {
		channel <- "Hello World5"
	}()
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// close(channel)

	fmt.Println(len(channel))
}
