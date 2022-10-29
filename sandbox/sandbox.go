package main

import (
	"fmt"
)

func main() {
	// for initialzation; condition ; post {
	// 	//body
	// }

	for i := 1; i < 5; i++ {
		if i == 1 {
			continue
		}
		fmt.Println(i)
		if i == 3 {
			fmt.Println("I am about to terminate")
			break
		}
	}

	// i := 1
	// for i < 5 {
	// 	fmt.Println(time.Now())
	// 	i++
	// }
}
