package main

import "fmt"

func main() {
	var n, hasil int
	fmt.Print("Masukkan batas angka n: ")
	fmt.Scan(&n)
	hasil = 0
	for i := 1; i <= n; i ++ {
		hasil += i
	}
	fmt.Println("Total penjumlahan dari 1 sampai", n, "adalah:", hasil)

}