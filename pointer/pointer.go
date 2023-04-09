package main

import "fmt"

func main() {
	a := 100         // Khởi tạo biến a có giá trị 100
	var pointer *int // Khai báo 1 biến con trỏ có kiểu dữ liệu là int

	pointer = &a // Biến pointer trỏ tới biến a

	fmt.Println(pointer) // In ra địa chỉ, vùng nhớ mà trỏ đến biến a

	fmt.Printf("%T", pointer) // In ra kiểu dữ liệu của pointer
}
