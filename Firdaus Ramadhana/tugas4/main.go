package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	fmt.Println("Hasil faktorialnya adalah: ", hasil)
}