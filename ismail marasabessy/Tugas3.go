package main

import "fmt"

func main() {
	var a, b, hasil int
	fmt.Print("Masukkan dua bilangan: ")
	fmt.Scan(&a, &b)
	hasil = 1
	for i := 1; i <= b; i++ {
		hasil *= a
	}
	fmt.Println("Hasil pemangkatan =", hasil)
}
