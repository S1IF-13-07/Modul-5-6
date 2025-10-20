package main

import "fmt"

func main() {
	var i, n, r, t int 
	fmt.Print("Masukkan banyaknya kerucut : ")
	fmt.Scan(&n)

	for i = 1 ; i <= n ; i++{
		fmt.Print("Masukkan jari-jari : ")
		fmt.Scan(&r)
		fmt.Print("Masukkan tinggi : ")
		fmt.Scan(&t)

		volume := (1.0 / 3.0) * 3.14 * float64(r * r * t)
		fmt.Println(volume)

	}
}