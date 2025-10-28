package main

import "fmt"

func main() {
	var n int
	var faktorial uint64 = 1
	fmt.Scan(&n)
	if n > 0 {
		for i := 1; i <= n; i++ {
			faktorial *= uint64(i)
		}
	}
	fmt.Println(faktorial)
}
