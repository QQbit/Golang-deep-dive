package main

import (
	"fmt"
)

func main() {
	// concept การทำงานพื้นฐานของ switch  สมุติว่ามี เคส if-else หลายๆอัน เราก็จะ simplify control flow ตรงนั่นให้ใช้ switch case
	x := 200
	switch x {
	case 1:
		fmt.Println("one")
		//ใน go ไม่จำเป็นต้องใช้ break ทุกครั้งที่ condition ถูก apply ก็จะทำ case นี้แล้วจบการทำงานเลยทันที
	case 2:
		fmt.Println("two")
		fallthrough //ถ้าต้องการไม่ให้ break
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}

	//switch สามารถใส่ expression ได้
	switch true { //true ตรงนี้สามารถ omit ได้
	case 1 <= x && x < 100:
		fmt.Println("deci")
	case 100 <= x && x < 1000:
		fmt.Println("centi")
	case 1000 <= x:
		fmt.Println("meter")
	default:
		fmt.Println("default")
	}

	// switch ใช้กับ string ได้ด้วย
	y := "z"
	switch y { //true ตรงนี้สามารถ omit ได้
	case "a":
		fmt.Println("A")
	case "b":
		fmt.Println("B")
	case "c", "d", "e", "f": //multiple case
		fmt.Println("c", "d", "e", "f")
		if y == "f" {
			break //สามารถมี break ได้ด้วย
		}
		fmt.Println("You are lucky")
	default:
		fmt.Println("DEFAULT")
	case "z":
	}
}
