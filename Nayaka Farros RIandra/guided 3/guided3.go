package main

import "fmt"

func main() {
	var a, b int
	var hasil int

	fmt.Print("masukkan bilangan pertama = ")
	fmt.Scan(&a)
	fmt.Print("masukkan bilangan kedua = ")
	fmt.Scan(&b)

	for i := 0; i < a; i++ {
		hasil += b
	}
	fmt.Print("hasil = ", hasil)
}