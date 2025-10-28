package main

import "fmt"

func main(){
	var n int
	var jariJari, tinggi, hasil float64
	const pi = 3.14159
	fmt.Print("Input banyak kerucut : ")
	fmt.Scan(&n)
	for i := 0; i < n ; i++{
		fmt.Scan(&jariJari)
		fmt.Scan(&tinggi)
		hasil = pi * jariJari * jariJari * tinggi / 3 
		fmt.Printf("Output Kerucut %d : %f\n",i+1, hasil)
	}
}