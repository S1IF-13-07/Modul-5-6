package main

import "fmt"

func main() {
	var a, b, hasil int
	fmt.Print("Masukan bilangan dan pangkat : ")
	fmt.Scan(&a, &b)

	hasil = 1
	for i := 1; i <= b; i++ {
		hasil *= a
	}
	fmt.Println("Hasil =", hasil)
}
