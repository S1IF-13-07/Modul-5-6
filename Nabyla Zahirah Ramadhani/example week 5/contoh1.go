package main

import "fmt"

func main (){
	var a, b int

	fmt.Print("Masukkan a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan b: ")
	fmt.Scan(&b)
	for i:=a; i<=b; i++{	
		fmt.Print(i)
	}
}