package main

import (
	"Laprak_Alpro_5/unguided"
	"fmt"
)

func main() {
	var a int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	hasil := unguided.Unguided1(a)
	fmt.Println("Hasil:", hasil)
}
