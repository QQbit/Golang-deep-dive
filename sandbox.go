package main

import (
	"fmt"
)

func main() {
	x := 5
	if (x <= 5) || false {
		fmt.Println("if condition")
	} else {
		fmt.Println("else condition")
	}

	// อีก feature หนึ่งของ go เราสามารถประกาศ local variable สำหรับ block ของ if - else ได้
	if v := x + 1; (x <= 5) || false {
		fmt.Println("if condition", v)
	} else {
		fmt.Println("else condition", v)
	}
}
