package main

import (
	"fmt"
	weight "weightx" //ทำการ rename package
	/**https://go.dev/ref/spec#Import_declarations
	The interpretation of the ImportPath is implementation-dependent but it is typically a substring of the full file name of the compiled package and may be relative to a repository of installed packages.
	ขึ้นอยู่กับเครื่องมือว่าจะให้ ความหมายของ ImportPath ตรงนี้ยังไง แต่โดยทั่วไปจะเป็น substring ของ fullname ของตัว compiled package นั่นอีกที่ หรือจะเป็น path ของ repository หรือ uri ของ repositiory
	**/)

func main() {
	k := weight.KG(1)
	fmt.Println(weight.KgToLb(k))
	/*
		weight.kgToLb(k) - kgToLb not exported by package weight
		go จะสามารถใช้งานแบบ public ได้ kgToLb (ตัวหน้าสุด) จะต้องเป็นตัวใหญ่ (capital letter) ถึงจะ export ได้
		ถ้าเป็นตัวเล็กเราจะไม่สามารถ export สามารถใช้ได้เฉพาะ package ของมันเท่านั่น
	*/
}
