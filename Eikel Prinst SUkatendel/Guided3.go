package main

import "fmt"

func main() {
	var a, b, hasil int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	for i := 1; i <= b; i++ {
		hasil += a
	}

	fmt.Println("Hasilnya:", hasil)
}