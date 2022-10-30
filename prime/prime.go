package prime

import (
	"fmt"
)

//var primes = [5]bool{false, false, true, true}

/*fmt.Println(prime.IsPrime(391232)) ถ้าเราต้องมีแบบนี้ เราต้องมา init แบบ line 3 หรือป่าว
แล้วถ้ามีสัก 9 แสนหละ เราต้อง init ทั้งหมดเลยไหม ไม่ต้องไม่จำเป็น มันจะมี func หนึ่งที่เรียกว่า init*/

var notPrimes [1_000_000]bool //ถ้าเราไม่ได้ทำการ initial มันทุกอย่างจะเป็น false

func init() { //จะไม่สามารถเรียกใช้โดยโปรแกรม แต่ว่ามันจะถูก invoke อัตโนมัติ เราไม่สามารถเรียกใช้ได้ตรงๆ ถ้าเรียกได้ตรงๆมันก็จะเป็น func ธรรมไปซะ
	//primes
	// algorithm to calculate prime https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
	fmt.Println("initialzation in Prime package")
	//sqrtLen := int(math.Ceil(math.Sqrt(float64(len(notPrimes)))))
	//for i := 2; i < sqrtLen; i++  แทนที่เราจะให้เช็คทุกค่า เราเช็คแค่ square roots ของมันก็พอ
	for i := 2; i < len(notPrimes); i++ {
		if notPrimes[i] {
			continue
		}
		notPrimes[i] = false
		for j := i * 2; j < len(notPrimes); j += i {
			notPrimes[j] = true
		}
	}
	fmt.Println("initialized")
}

/*
	สรุป ใน func init จะถูก execute แค่ครั้งเดียว และ ก็ทำการ initalize variable ต่างๆ ในระดับ package level
	ทุกครั้งที่เรามาใช้ function ที่เรา export ไปให้ client คนอื่นเขามาใช้ก็ทำการ access local variable ของเราคือ notPrimes ใน line 34
*/

func IsPrime(x int) bool {
	return !notPrimes[x]
}
