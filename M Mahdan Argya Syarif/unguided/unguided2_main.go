package main

import (
	"Laprak_Alpro_5/unguided"
	"fmt"
)

func main() {
	var r, t float64
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&r, &t)
		hasil := unguided.Unguided2(n, r, t)
		fmt.Println("Hasil:", hasil)
	}
}
