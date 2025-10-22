package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var r, t float64
	
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		fmt.Scan(&r, &t)
		volume := (1.0 / 3.0) * math.Pi * r * r * t
		fmt.Println(volume)
	}
}