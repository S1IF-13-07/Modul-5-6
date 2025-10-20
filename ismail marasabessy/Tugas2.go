package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah kerucut: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var r, t float64
		fmt.Scan(&r, &t)
		volume := (1.0 / 3.0) * math.Pi * r * r * t
		fmt.Println(volume)
	}
}
