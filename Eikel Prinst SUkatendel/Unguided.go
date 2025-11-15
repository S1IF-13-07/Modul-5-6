package main

import "fmt"

func main() {
	var n, hasil int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		hasil += i
	}

	fmt.Println("Hasil:", hasil)
}