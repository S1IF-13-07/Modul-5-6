package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var pangkat int
	var hasil int

	fmt.Print("Masukkan angka pertama : ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai pangkat : ")
	fmt.Scanln(&pangkat)

	for i := 1; i <= pangkat; i++ {
		hasil = int(math.Pow(float64(a), float64(i)))

	}
	fmt.Printf("Hasilnya adalah %d\n", hasil)
}
