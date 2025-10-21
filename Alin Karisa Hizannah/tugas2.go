package main

import (
	"fmt"
	"math"

)

func main() {
	var n int
	fmt.Print("Masukkan Banyak kerucut/banyak baris (n): ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var r, t float64
		fmt.Print("Masukkan jari-jari, tinggi: ")
		fmt.Scan(&r, &t)
		volume := (1.0 / 3.0) * math.Pi * r * r * t
		fmt.Println("Volume kerucut baris ke", i, "adalah", volume)
	}
}
