package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var r, t, volume float64
	
	fmt.Scan(&n)

	for j := 0; j < n; j++ {
		fmt.Scan(&r, &t) 

		volume = (1.0 / 3.0) * math.Pi * math.Pow(r, 2) * t
		
		fmt.Println(volume)
	}
}