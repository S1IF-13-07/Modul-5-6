package main

import "fmt"

func main() {
	var a, b, c int
	var hasil int
	fmt.Print("masukan angka pertama: ")
	fmt.Scan(&b)
	fmt.Print("masukan angka kedua: ")
	fmt.Scan(&c)
	hasil = 0 
		for a = 1; a <= c; a+=1 { 
		hasil = hasil + b
	}
	fmt.Println(hasil)
}