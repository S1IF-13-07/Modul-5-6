package main

import "fmt"

func main() {
	var a, b int
	var hasil int = 1
	fmt.Print("Input bilangan dan pangkat: ")
	fmt.Scan(&a, &b)

	for i := 0; i < b; i++ {
		hasil = hasil * a
	}
	fmt.Println("Hasil:", hasil)
}