package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan bilangan pertama (a): ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan kedua (b): ")
	fmt.Scan(&b)

	hasil := 1
	for i := 1; i <= b; i++ {
		hasil *= a
	}

	fmt.Printf("%d pangkat %d = %d\n", a, b, hasil)
}
