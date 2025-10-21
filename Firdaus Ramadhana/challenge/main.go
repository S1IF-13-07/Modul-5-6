package main 

import "fmt"

func main() {
	var nilai int
	fmt.Print("masukkan nilai masbro: ")
	fmt.Scan(&nilai)
	for i := 1; i <= nilai; i++ {
		for s := 1 ; s <= nilai-i ; s++ {
			fmt.Print(" ")
		} 

		for j := 1 ; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println() 
	} 
}