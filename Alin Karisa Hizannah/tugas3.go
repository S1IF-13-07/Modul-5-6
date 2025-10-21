package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan dua bilangan (basis dan pangkat): ")
	fmt.Scan(&a, &b)

	hasil := 1
	for i := 0; i < b; i++ {
		hasil *= a
	}

	fmt.Println("Hasil:", hasil)
}
