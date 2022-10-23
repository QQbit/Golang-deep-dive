package main

import "fmt"

// การสร้าง Alias ไม่ได้มีเฉพาะการ copy pointer เท่านั่น
// ยังมีทุกครั้งที่เราทำการ copy reference type เช่น map , slices , channel, struct , array ,interface
func double(x *int) { // x ตรงนี้ คือ Alias ของ x
	*x = *x * 2
}
func main() {
	x := 2
	double(&x)
	fmt.Println(x) // 4
}
