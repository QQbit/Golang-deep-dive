package main

import "fmt"

func main() {
	// Go variable declearation
	// 1. var name type = expression
	var x int = 4
	fmt.Println(x)

	// 2. var name = expression แบบ omit type
	var y = 5 //มันก็จะทำการ deduce type
	fmt.Println(y)

	var a = 5
	var b = 5.0
	var c = "string"
	fmt.Printf("%T\n", a) //int
	fmt.Printf("%T\n", b) //float64
	fmt.Printf("%T\n", c) //string

	// 3. var name type แบบ omit expression
	var d int
	var e string
	fmt.Println(d) // zero value of int is 0
	fmt.Println(e) // zero value of string is "" (empty string)

	//ถ้าไม่ใช่ primitive type เป็นแบบ reference type
	var f []string        //ยกตัวอย่างเป็น slices of string
	fmt.Println(f == nil) //true  zero value of slices is nil
	fmt.Println(f)        // []

	//Go สามารถประกาศตัวแปรหลายๆตัวได้ ใน statement เดียวกัน
	var g, h, i int
	fmt.Printf("%T %T %T\n", g, h, i)
	fmt.Printf("%v %v %v\n", g, h, i)

	var j, l, m = 123, "hello", 4.5
	fmt.Printf("%T %T %T\n", j, l, m)
	fmt.Printf("%v %v %v\n", j, l, m)
}
