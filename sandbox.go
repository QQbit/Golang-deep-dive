package main

import "fmt"

test := 123  // อีก limit หนึ่งของ short declaration 
//syntax error: non-declaration statement outside function body เพราะ line 5 ไม่ใช่ declaration statement แต่มันเป็น short declaration 
var test2 = 1 //ต้องใช้แบบนี้ถึงได้

func main() {
	// Go short variable declaration
	// 1. var name = expression
	var a = 4
	fmt.Println(a)
	b := 4
	fmt.Println(b)
	x, y := 1, 2
	fmt.Println(x, y)
	//x, y := y, x no new variables on left side of :=
	//ทำไม่ได้ กฏของ go คือต้องมีตัวใดตัวหนึ่งเป็น varible ใหม่
	x, y1 := y, x
	fmt.Println(x, y1)

	//มาดูกันว่าทำไม เราถึง short declaration 2 รอบไม่ได้
	i := 1 //var x = 1 
	i := 2 //var x = 2  - no new variables on left side of :=  ก็คือ variable เราประกาศไว้แล้ว เราประกาศอีกมันก็ผิดกฏ 
	fmt.Println(i)
}
