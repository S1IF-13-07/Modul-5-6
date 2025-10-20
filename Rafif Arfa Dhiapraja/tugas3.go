package main

import "fmt"

func main() {
	var i, angka, pangkat, hasil int
	fmt.Print("Masukkan angka : ")
	fmt.Scan(&angka)
	fmt.Print("Masukkan Pangkat : ")
	fmt.Scan(&pangkat)

	hasil = 1

	for i = 1 ; i <= pangkat ; i++{
		hasil *= angka 

	}

	fmt.Println("hasil pemangkatan :", hasil)	
}