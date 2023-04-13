package main

import "fmt"

func main() {
	// zero value
	var pointer *int // Khai báo con trỏ với từ khóa new thì sẽ không bị nil

	a := 100

	pointer = &a

	fmt.Println(a)

	*pointer = 999 // a:= 999

	fmt.Println(pointer)
	fmt.Println(a)
}
