package main

import "fmt"
func main(){
	var a,b int
	fmt.Print("Masukan Nilai Pertama : ")
	fmt.Scan(&a)

	fmt.Print("Masukan Nilai Kedua : ")
	fmt.Scan(&b)
	var j int
	for j = a; j <=b; j+=1{
	fmt.Print(j, " ")}
}