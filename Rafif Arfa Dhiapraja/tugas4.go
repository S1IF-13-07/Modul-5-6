package main

import "fmt"

func main() {
	var n, i, hasil int
	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&n)

	hasil = 1

	for i = 1 ; i <= n ; i++{
		hasil *= i
	}

	fmt.Println(hasil)
}