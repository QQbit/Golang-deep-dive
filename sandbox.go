package main

import (
	"fmt"
)

type KG float64
type LB float64

func (lb LB) toKg() KG {
	return KG(lb / 2.2046226218)
}

func (kg KG) toString() string {
	return fmt.Sprintf("%v kg", kg)
}

func main() {
	k := KG(3)
	b := LB(3)
	fmt.Println(k > b.toKg())
	fmt.Println(k.toString())
}
