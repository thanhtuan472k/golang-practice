package main

import "fmt"

func main() {
	b := 123
	p2 := new(int) // <=> var p2 *int

	p2 = &b

	fmt.Println()
	fmt.Println(p2)
	fmt.Printf("%T", p2)
}
