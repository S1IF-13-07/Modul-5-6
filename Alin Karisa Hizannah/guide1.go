package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan bilangan a : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan b : ")
	fmt.Scan(&b)
	for i := a; i <= b; i+=1 {
		fmt.Print(i, " ")
	}
}
