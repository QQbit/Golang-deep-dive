package main

import (
	"fmt"

	"github.com/QQbit/weight"
	gl "gitlab.com/kim0x/weight"
)

func main() {
	k := weight.KG(1)
	k2 := gl.KG(.15)
	fmt.Println(gl.KgToLb(k + k2)) //invalid operation: k + k2 (mismatched types "github.com/QQbit/weight".KG and "gitlab.com/kim0x/weight".KG)
	/*
		สรุป ถึงแม้ว่าโค้ดข้างในจะเหมือนกัน ถ้า package ที่ import มันต่างกัน มันก็จะไม่สามารถใช้งานร่วมกันได้
		ในกรณีที่เป็น type declearation เท่านั่น ถ้าเป็น constant ไม่มีปัญหา เพราะว่า constant จาก math.PI หรือจากอะไรก็ตาม สามารถใช้ร่วมกันได้
	*/
}
