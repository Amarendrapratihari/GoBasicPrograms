/*package main

import (
	"fmt"
)

func main() {
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4] //[2:4]--->>>>>>>>>[Array Index 2 and 4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
}*/

package main

import "fmt"

func main() {

	a := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	b := a[:]   //slice of all elements
	c := a[3:]  //slice from 4th element to end
	d := a[:6]  //slice upto 6th element
	e := a[3:6] //slice the first 4th, 5th, 6th element
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
