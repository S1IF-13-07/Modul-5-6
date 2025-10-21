package main

import "fmt"

func main() {
	var basis, pangkat int

	fmt.Print("Masukkan bilangan bulat =  ")
	fmt.Scan(&basis)

	fmt.Print("Masukkan bilangan eksponen = ")
	fmt.Scan(&pangkat)

	var hasil int = 1
	basisHitung := basis
	pangkatHitung := pangkat

	if pangkat < 0 {
		basisHitung = 1 / basis 
		pangkatHitung = pangkat  
	}
	for i := 0; i < pangkatHitung; i++ {
		hasil = hasil * basisHitung
	}
	fmt.Print("Hasil dari pangkat adalah = ", hasil)
}