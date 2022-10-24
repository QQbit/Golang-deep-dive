package main

import (
	"fmt"
	"math"
)

// type YourTypeName underlyingType
type KG float64
type LB float64

//1 kilogram (kg) = 2.2046226218 pounds (lbs)
/* comparison
[KG - float64] == [LB - float64]  false
[KG - float64] == [float64]		  false
[LB - float64] == [string]		  false
*/
func main() {
	k := KG(3)
	b := LB(3)
	fmt.Println(k == b)          // mismatched types KG and LB
	fmt.Println(k == float64(3)) //mismatched types KG and float64  - type explicitly by a constant coversion

	/*
		https://go.dev/ref/spec#Constants
		A constant may be given a type explicitly by a constant declaration or conversion,
		or implicitly when used in a variable declaration or an assignment statement or as an operand in an expression.
	*/

	fmt.Println(k == 3) // true ทำไมมันเทียบได้ เพราะ 3 คือ constant
	// และ 3 ก็จะถูก  assign type แบบ implicitly เพราะเป็น  as an operand in an expression. ก็จะ represent KG ได้

	fmt.Println(k == "3") //mismatched types KG and untyped string เพราะ string ไม่สามารถ represent float64 ไม่สามารถ assign ตรงนี้ให้เป็น KG ได้

	result := k + 3            //implicitly when used in  an assignment statement
	fmt.Printf("%T/n", result) //main.KG

	//constant มีอะไรอีกบ้าง
	/*
		// Mathematical constants.
		const (
			Pi  = 3.14159265358979323846264338327950288419716939937510582097494459 // https://oeis.org/A000796
		)
	*/
	MyResult := k + math.Pi //ไม่ได้บอกว่า Pi คือ float64 หรือ float32 บอกแค่ว่า Pi ตรงนี้คือ constant (line 42)
	//คนที่ทำภาษา go เขาไม่รู้หรอกว่า วันดีคืนดี เราจะมาประกาศเป็น KG ถ้าเราเอา Pi มา + กับ KG แล้วจะเกิดอะไรขึ้น เขาไม่รู้หรอก แต่พอ spec บอกมาแล้ว
	//เขาก็จะ convert math.Pi ไปเป็น KG
	fmt.Printf("%T/n", MyResult) //main.KG
	fmt.Println(MyResult)

	/* comparison เป็นไปได้ทางเดียวคือ
	[KG - float64] == [KG - float64]  true
	[LB - float64] == [LB - float64]  true
	*/
}
