package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var r, t float64
		fmt.Scan(&r, &t)

		volume := (1.0 / 3.0) * math.Pi * math.Pow(r, 2) * t

		fmt.Println(volume)
	}
}
