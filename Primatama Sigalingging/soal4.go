package main 

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	fmt.Println("Hasil faktorial dari", n, "adalah:", hasil)
}