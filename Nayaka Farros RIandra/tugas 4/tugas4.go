package main

import "fmt"

func main() {
	var n int
	
	fmt.Print("masukkan bilangan bulat = ")
	fmt.Scan(&n)

	total := 1
	for i := 1; i <= n; i+=1 {
		total *= i
	}
	fmt.Println("Jumlah = ", total)
}