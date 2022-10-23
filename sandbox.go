package main

import "fmt"

func newIntPointer() *int {
	var x int
	return &x
}
func main() {
	fmt.Println(*newIntPointer())                   //0 zero value
	fmt.Println(*new(int))                          //0 zero value โดย new() เป็น built-in function  จะทำงานเหมือน function newIntPointer ทุกอย่าง
	fmt.Println(new(string) == new(string))         //false
	fmt.Println(newIntPointer() == newIntPointer()) //false
}
