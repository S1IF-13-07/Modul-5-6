package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("masukkan bilangan pertama = ")
	fmt.Scan(&a)
	fmt.Print("masukkan bilangan kedua = ")
	fmt.Scan(&b)

	for i  := a; i  <= b; i ++ {
		fmt.Print(i)
	}
}