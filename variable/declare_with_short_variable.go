package main

import "fmt"

func main() {
	// Khi khai báo biến bằng cách := thì Golang sẽ tự động suy ra kiểu dữ liệu dựa trên giá trị được gán
	// Không thể khai báo thêm kiểu dữ liệu cho cách khai báo :=
	// Khai báo bằng := thì biến chỉ có giá trị bên trong hàm (giống local scope bên JS)
	// Khai báo bằng := thì không thể thay đổi giá trị của biến
	a := 5

	fmt.Println(a)

}
