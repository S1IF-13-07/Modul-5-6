package main

import "fmt"

func main() {

	var pertama, kedua int
	fmt.Print("Masukan Nilai Pertama : ")
	fmt.Scan(&pertama)

	fmt.Print("Masukan Nilai Kedua : ")
	fmt.Scan(&kedua)

	for i := pertama; i <= kedua; i++ {
		fmt.Print(i, " ")
	}

}
