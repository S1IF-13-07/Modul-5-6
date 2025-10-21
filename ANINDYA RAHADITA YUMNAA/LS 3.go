package main

import "fmt"

func main() {
	var basis, eksponen int
	var hasil int = 1

	fmt.Scan(&basis, &eksponen)

	for i := 0; i < eksponen; i++ {
		hasil = hasil * basis
	}

	fmt.Println(hasil)
}