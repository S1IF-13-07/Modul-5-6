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
		fmt.Printf("Masukkan jari-jari dan tinggi kerucut ke-%d: ", i)
		fmt.Scan(&r, &t)

		volume := (1.0 / 3.0) * math.Pi * math.Pow(r, 2) * t
		fmt.Printf("Volume kerucut ke-%d = %.14f\n", i, volume)
	}
}
