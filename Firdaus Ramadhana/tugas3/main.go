package main 

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan 2 nilai bulat positif: ")
	fmt.Scan(&a, &b)
	hasil := 1
	for i := 1; i <= b; i++ {
		hasil *= a
	}
	fmt.Printf("Hasil perpangkatan dari bilangan %d dipangkat dengan %d adalah %d", a, b, hasil)
}