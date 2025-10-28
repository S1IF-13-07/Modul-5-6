package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var b, c, volume float64
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		fmt.Scan(&b, &c)
		volume = (1.0 / 3.0) * math.Pi * math.Pow(b, 2) * c
		fmt.Println("hasil: ", volume)
	}
}
