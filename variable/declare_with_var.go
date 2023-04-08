package main

import "fmt"

var b = 6

func main() {
	// Khai báo với var
	//var a = 10 // Khai báo như này cũng được
	var a int64 = 10 // Khai báo biến với var có thể khai báo thêm kiểu dữ liệu cho giá trị cần gán
	fmt.Println("a =", a)

	a = 15
	fmt.Println("a =", a) // Khai báo var có thể ga lại giá trị khác cho biến đó

	fmt.Println("b =", b) // Phạm vi khai báo var là global scope, đối với khai báo := thì không thể làm được như vậy
}
