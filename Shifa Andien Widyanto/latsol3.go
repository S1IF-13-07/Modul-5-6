package main

import "fmt"

func main(){
	var a, b, hasil int
	fmt.Print("Masukan bilangan bulat a : ")
	fmt.Scan(&a)
	fmt.Print("Masukan bilangan bulat b : ")
	fmt.Scan(&b)
	for i := 1; i <= b; i++ {
		hasil += a
	}
	fmt.Println(hasil)
}