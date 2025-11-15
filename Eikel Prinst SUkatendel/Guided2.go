package main

import "fmt"

func main() {
	var n, alas, tinggi int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print("Masukkan alas dan tinggi: ")
		fmt.Scan(&alas, &tinggi)

		luas := 0.5 * float64(alas*tinggi)
		fmt.Println("Luasnya:", luas)
	}
}