package main

import "fmt"

func main() {
    var n int
	fmt.Print("Masukan bil n: ")
    fmt.Scan(&n)
    hasil := 1
    	for i := 1; i <= n; i++ {
        hasil *= i
    }
    fmt.Println(hasil)
}
