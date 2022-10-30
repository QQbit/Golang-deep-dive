package main

import (
	"fmt"
	"math/rand"
	"prime"
)

func main() {

	for i := 0; i < 100; i++ {
		x := rand.Intn(1000000)
		fmt.Printf("%d , %t\n", x, prime.IsPrime(x))
	}

	fmt.Println(prime.IsPrime(2))
	fmt.Println(prime.IsPrime(3))
	fmt.Println(prime.IsPrime(4))
	fmt.Println(prime.IsPrime(491023)) //ถ้าเกิดว่าทุกครั้งที่เราเรียก isPrime แล้วมาต้องคำนวณใหม่หมด 1,000,000 ครั้ง เราจะต้องมาเช็คว่าตรงนั่นเป็น prime หรือเปล่า มันช้า
}
