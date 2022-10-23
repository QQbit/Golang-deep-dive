package main

import "fmt"

// address จะเปลี่ยนทุกรอบที่รัน เพราะขึ้นอยู่กับ os จะจัดการว่าให้โปรแกรมของเราไปรันใน section ไหนของ memory
func main() {
	x := 70
	var p *int
	fmt.Println("&x", &x) //represent address in memory
	fmt.Println("&p", &p) //เป็น nil เพราะไม่มีค่า แต่จะต้องมี address ที่ associate มันใน main memory
	p = &x
	fmt.Println("p", p)           //0x14000126008
	fmt.Println("vale of *p", *p) //dereference pointer ก็จะได้ value ของ x ซึ้งก็คือ 70
	fmt.Println("*(&x)", *&x)     //70
	/**
	-------------------------------------------------
	| variable | address   		|   value   		|
	-------------------------------------------------
	| 	 x	   | 0x14000126008	| 	 70				|
	-------------------------------------------------
	| 	 p	   | 0x14000120018	| 	 0x14000126008	|
	-------------------------------------------------


	**/
}
