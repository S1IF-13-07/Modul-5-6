package main

import "fmt"

func main() {
	var n int
	var hasil int = 1

	
	fmt.Print("Masukkan bilangan bulat non-negatif: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		hasil *= i
	}

	
	fmt.Println("Hasil faktorial dari", n, "adalah:", hasil)
}
