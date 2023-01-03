package main

import (
	"fmt"
)

func mainr() {
	arr1 := [5]string{"Volvo", "BMW", "Ford", "Mazda", "Tata"}
	arr2 := [...]int{1, 2, 3, 4}

	//Find the length of an Array by using-->>>>> "len()" <<<<<---.
	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
}
