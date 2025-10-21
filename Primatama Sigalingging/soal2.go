package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var r, t float64
	fmt.Print("Berapa banyak kerucut yang ingin dihitung? : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i ++ {
		fmt.Print("Masukkan jari jari dan tinggi kerucut ke-", i, ": ")
		fmt.Scan(&r, &t)

		volume := (1.0 / 3.0) *math.Pi * r * r * t
		fmt.Println("Volume kerucut ke-", i, "adalah", volume)
	}
}