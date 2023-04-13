package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}

	var pointer *[3]int

	pointer = &array
	fmt.Println(pointer)      // &[1, 2, 3]
	fmt.Printf("%T", pointer) // *[3]int
}
