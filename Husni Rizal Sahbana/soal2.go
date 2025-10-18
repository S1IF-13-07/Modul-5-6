package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var r, t, volume float64

	fmt.Print("Masukan Jumlah Kerucut : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print("Masukan jari-jari dan tinggi kerucut ke-", i, ": ")
		fmt.Scan(&r, &t)

		volume = (1.0 / 3.0) * math.Pi * r * r * t
		fmt.Printf("Volume kerucut ke-%d = %.14f\n", i, volume)
	}
}
