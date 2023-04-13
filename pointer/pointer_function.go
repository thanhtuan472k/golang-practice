package main

import "fmt"

func applyPointer(pointer *int) {
	*pointer = 777
	//pointer = 777 => không sử dụng được, báo lỗi
}

func main() {
	c := 30
	var pointer *int = &c
	applyPointer(pointer)

	fmt.Println(c)

}
