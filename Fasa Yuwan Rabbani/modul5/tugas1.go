package main

import "fmt"

func main() {
	var a int
	var sum int
	fmt.Print("Masukkan sebuah angka bilangan bulat positif : ")
	fmt.Scanln(&a)

	for i := 1; i <= a; i++ {
		sum += i
	}
	fmt.Printf("Hasil dari %d adalah %d\n", a, sum)
}
