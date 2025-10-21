package main

import "fmt"

func main() {
	var j, alas, tinggi, n int
	var luas float64

	fmt.Print("Masukkan banyak segitiga (n): ")
	fmt.Scan(&n)

	for j = 1; j <= n; j++ {
		fmt.Printf("Masukkan alas dan tinggi segitiga ke-%d: ", j)
		fmt.Scan(&alas, &tinggi)
		luas = 0.5 * float64(alas*tinggi)
		fmt.Printf("Luas segitiga ke-%d adalah %.1f\n", j, luas)
	}
}
