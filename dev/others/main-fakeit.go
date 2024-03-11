package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {
	gofakeit.Seed(0)

	fmt.Println(gofakeit.IPv4Address())
	// fmt.Println(gofakeit.IPv4Address())
	// fmt.Println(gofakeit.IPv4Address())
	// fmt.Println(gofakeit.IPv4Address())
	// fmt.Println(gofakeit.IPv4Address())
	// fmt.Println(gofakeit.IPv4Address())
}
