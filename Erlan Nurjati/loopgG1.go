package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("masukan angka pertama: ")
	fmt.Scan(&a)
	fmt.Print("masukan angka kedua: ")
	fmt.Scan(&b)
	for i := a; i<=b; i++ {
		fmt.Print(i)
	}
}
