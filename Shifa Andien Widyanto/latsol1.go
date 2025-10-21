package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukan bilangan bulat a : ")
	fmt.Scan(&a)
	fmt.Print("Masukan bilangan bulat b  : ")
	fmt.Scan(&b)
	for i :=a; i<=b; i++ {
		fmt.Print(i)
	}
}