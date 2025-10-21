package main

import "fmt"

func main() {
	var n, hasil int
	fmt.Print("Masukan nilai n : ")
	fmt.Scan(&n)

	hasil = 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	fmt.Println("Hasil faktorial =", hasil)
}
