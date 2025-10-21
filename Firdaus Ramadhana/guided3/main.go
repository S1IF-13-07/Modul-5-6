package main

import "fmt"

func main() {
	var a, b int
	var hasil int
	fmt.Print("masukkan nilai a nya mas: ")
	fmt.Scan(&a)
	fmt.Print("masukkan nilai b nya mas: ")
	fmt.Scan(&b)
	for i := 1; i <= b; i++ {
		hasil = hasil + a
	}
fmt.Print("Hasilnya mas: ", hasil)
}