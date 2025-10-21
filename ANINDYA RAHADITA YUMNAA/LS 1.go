package main

import "fmt"

func main() {
	var n int
	var total int = 0
	
	fmt.Scan(&n)

	for j := 1; j <= n; j++ {
		total += j
	}
	
	fmt.Println(total)
}