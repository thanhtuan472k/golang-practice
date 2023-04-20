package main

import "fmt"

func log() {
	fmt.Println("aaa")
}

func main() {
	defer fmt.Println("World")
	defer fmt.Println("World1")

	fmt.Println("Hello")

	log()
}
