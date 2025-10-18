package main

import "fmt"

func main() {
	var bil_1, bil_2, jumlah int
	fmt.Print("Masukan Bilangan Pertama : ")
	fmt.Scan(&bil_1)

	fmt.Print("Masukan Bilangan Kedua : ")
	fmt.Scan(&bil_2)

	jumlah = 1
	for i := 1; i <= bil_2; i++ {
		jumlah *= bil_1
	}
	fmt.Println(jumlah)
}
