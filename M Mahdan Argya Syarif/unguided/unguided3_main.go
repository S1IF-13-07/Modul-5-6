package main

import (
	"Laprak_Alpro_5/unguided"
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a dan b: ")
	fmt.Scan(&a, &b)
	hasil := unguided.Unguided3(a, b)
	fmt.Println("Hasil:", hasil)
}
