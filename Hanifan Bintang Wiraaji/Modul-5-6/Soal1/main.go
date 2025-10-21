package main

import "fmt"

func main() {
	var a int
	var hasil int = 0
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		hasil = hasil + i
	}
	fmt.Print(hasil)
}