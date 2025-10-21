package main

import (
	"Laprak_Alpro_5/unguided"
	"fmt"
)

func main() {
	var a int
	fmt.Println("Masukan angka: ")
	fmt.Scan(&a)
	result1 := unguided.Unguided4(a)
	fmt.Println(result1)
}
