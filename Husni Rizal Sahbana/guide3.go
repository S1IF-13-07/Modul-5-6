package main

import "fmt"

func main() {
	var pertama, kedua, hasil int

	fmt.Print("Masukan bilangan 1 : ")
	fmt.Scan(&pertama)

	fmt.Print("Masukan bilnagan 2 : ")
	fmt.Scan(&kedua)

	for i := 1; i <= kedua; i++ {
		hasil += pertama
	}

	fmt.Print(hasil)
}
