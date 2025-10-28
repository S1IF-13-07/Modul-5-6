package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input n : ")
	fmt.Scan(&n)
	if n == 0 {
		fmt.Print("Output : ", 1)
	} else if n > 0 {
		for i := n; i > 1; i-- {
			n *= i - 1
		}
		fmt.Print("Output : ", n)
	}
}