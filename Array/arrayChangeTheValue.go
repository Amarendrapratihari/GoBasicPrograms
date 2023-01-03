package main

import (
	"fmt"
)

func main() {
	prices := [4]int{10, 20, 30, 40}

	prices[2] = 50
	fmt.Println(prices)
}
