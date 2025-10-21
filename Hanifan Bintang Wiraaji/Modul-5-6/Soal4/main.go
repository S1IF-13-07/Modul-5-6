package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var hasil = a
	for i := ( a - 1 ); i > 1; i-- {
		hasil = hasil * i
	}
	fmt.Print(hasil)
}
