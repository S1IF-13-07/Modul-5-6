package main

import "fmt"

func main() {
	var n int
	var r, t, volume float64
	const phi = 3.141592653589793

	fmt.Print("Masukan jumlah kerucut : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print("Masukan jari-jari dan tinggi : ")
		fmt.Scan(&r, &t)
		volume = (1.0 / 3.0) * phi * r * r * t
		fmt.Println("Volume kerucut ke-", i, ":", volume)
	}
}
