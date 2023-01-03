package main

import (
	"fmt"
)

// Declara a Function.
func myMessage() {
	fmt.Println("I just got executed!")
}
func sumData() {
	fmt.Println(5 + 9)
}

func main() {
	myMessage() // call the function
	sumData()   // call the function
}
