package main

import "fmt"

func fp() *int {
	//ถ้าเป็นภาษาอื่น local variable ถ้าเกิด function นี้จบการทำงาน function ใน stack memory ตรงนั่น ก็จะถูกลบหมด
	//แต่ว่า go มี Garbage Collector ที่ค่อยข้างฉลาดแล้วก็รู้ว่า function นี้ return address ของ local variable มานะ
	x := 4
	return &x
}
func main() {
	x := fp()
	fmt.Printf("%T\n", x)  // *int pointer of int
	fmt.Printf("%d\n", *x) // 4

	a := fp()
	fmt.Println(a == fp()) // false ทุกครั้งที่เรียกใช้ fp() มันก็จะสร้าง stack ใน memory แล้ว return ตัว local variable อันใหม่มาให้เราเสมอ
}
