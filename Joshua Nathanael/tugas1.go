package main

import "fmt"

func main() {
	var n, jumlah int
	fmt.Print("Masukan nilai n : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		jumlah += i
	}
	fmt.Println("Akan mengeluarkan hasil ", n, "adalah :", jumlah)
}