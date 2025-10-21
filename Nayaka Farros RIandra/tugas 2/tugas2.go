package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukkan jumlah kerucut = ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var r, t float64
		fmt.Print("masukkan panjang alas dan tinggi kerucut dari kerucut ke-", i, "= ")
		fmt.Scan(&r, &t)
		volume := (1.0 / 3.0) * 3.14 * r * r * t
		fmt.Printf("volume kerucut ke-%d adalah: %.2f\n", i, volume)
	}
}