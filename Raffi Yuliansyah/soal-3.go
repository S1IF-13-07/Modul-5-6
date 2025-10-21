package main

import "fmt"

func main(){
	var basis, pangkat, pengali int
	fmt.Print("Input b : ")
	fmt.Scan(&basis)
	fmt.Print("Input p : ")
	fmt.Scan(&pangkat)
	pengali = basis
	for i := 1; i < pangkat; i++ {
		basis *= pengali
	}
	fmt.Printf("Hasil : %d", basis)
}