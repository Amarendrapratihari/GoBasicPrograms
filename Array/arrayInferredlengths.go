package main

import (
	"fmt"
)

func main() {
	// here length is inferred
	var arr1 = [...]int{1, 2, 3, 9}
	arr2 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}
