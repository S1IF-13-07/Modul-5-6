package main

import "fmt"

func main() {
	var a int
	var factorial int = 1

	fmt.Print("Masukkan angka : ")
	fmt.Scanln(&a)

	for i := 1; i <= a; i++ {
		factorial *= i
	}
	fmt.Printf("Hasil faktorial dari %d adalah %d\n", a, factorial)
}
