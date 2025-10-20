package main

import "fmt"

func main () {
	var bilangan, jumlah int
	fmt.Print("Masukkan sebuah bilangan : ")
	fmt.Scan(&bilangan)

	for i := 1 ; i <= bilangan ; i++{
		jumlah += i

	}

	fmt.Println(jumlah)

}