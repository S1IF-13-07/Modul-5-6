package main

import "fmt"

func main() {
	var a, b int
	var hasil int = 1

	
	fmt.Print("Masukkan dua bilangan bulat positif basis dan pangkat: ")
	fmt.Scan(&a, &b)

	
	for i := 1; i <= b; i++ {
		hasil *= a
	}

	
	fmt.Println("Hasil:", hasil)
}
