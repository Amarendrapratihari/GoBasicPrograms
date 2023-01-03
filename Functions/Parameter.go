package main

import (
	"fmt"
)

func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

func main() {
	familyName("Hari")
	familyName("Swagatam")
	familyName("Ajay")
}
