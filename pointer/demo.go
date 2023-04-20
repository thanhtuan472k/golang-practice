package main

import "fmt"

func main() {
	var a = 100
	p := &a

	fmt.Println("&a", &a)

	fmt.Println(p)
}
