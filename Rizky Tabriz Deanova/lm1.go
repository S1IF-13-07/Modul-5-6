package main

import "fmt"

func main() {
	var a int
	var total int = 0
	fmt.Scan(&a)
	for b := 1; b <= a; b++ {
		total += b
	}
	fmt.Println(total)
}
