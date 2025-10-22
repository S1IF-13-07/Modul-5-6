package main

import "fmt"

func main() {
	var n int
	const pi = 3.141592653589793

	fmt.Print("Masukkan jumlah kerucut: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var r, t float64
		fmt.Print("Masukkan jari-jari dan tinggi kerucut ke-", i, ": ")
		fmt.Scan(&r, &t)

		
		volume := (1.0 / 3.0) * pi * r * r * t

		fmt.Println(volume)
	}
}
