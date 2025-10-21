package main

import "fmt"

func main() {
	var j, v1, v2,hasil int
	fmt.Scan (&v1, &v2)
	for j = 1; j <= v2; j+=1 {
		hasil += v1
	}
	fmt.Println(hasil)
}
