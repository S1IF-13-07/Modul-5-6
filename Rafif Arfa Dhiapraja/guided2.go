package main

import "fmt"

func main() {
	var i, alas, tinggi, n int
	var Luas float64

	fmt.Scan(&n)

	for i = 1; i <= n; i+=1 {
		fmt.Scan(&alas, &tinggi)
		Luas = 0.5 * float64(alas * tinggi)
		fmt.Println(Luas)
	}
}