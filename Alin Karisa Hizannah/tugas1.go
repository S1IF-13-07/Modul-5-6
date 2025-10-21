package main

import "fmt"

func main() {
	var n, hasil int
	fmt.Print("Masukkan Bilangan Positif : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		hasil += i
	}

	fmt.Println("Hasil Penjumlahan 1 sampai", n, "adalah:", hasil)
}
