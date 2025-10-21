package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var hasil = a
	for i := 1; i < b; i++ {
		hasil = hasil * a
	}
	fmt.Print(hasil)
}
