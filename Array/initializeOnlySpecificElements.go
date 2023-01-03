package main

import (
	"fmt"
)

func main() {
	//1:10 means: assign 10 to array index 1 (second element).
	//2:40 means: assign 40 to array index 2 (third element).
	arr1 := [5]int{1: 10, 2: 40}

	fmt.Println(arr1)
}
