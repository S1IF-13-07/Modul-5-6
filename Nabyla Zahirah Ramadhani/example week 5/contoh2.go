package main

import "fmt"

func main (){
	var a, b, hasil int
	fmt.Print("Masukkan a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan b: ")
	fmt.Scan(&b)
	for i:=1; i<=b; i++{
		hasil +=a
	}
	fmt.Print("Hasil perkalian: ", hasil)
}