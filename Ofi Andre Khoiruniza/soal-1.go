package main 

import "fmt"

func main(){
	var bilangan, hasil int
	fmt.Print("Input bilangan : ")
	fmt.Scan(&bilangan)
	for i := 1; i <= bilangan; i++{
		hasil += i
	}
	fmt.Printf("Output jumlah deret bilangan : %d", hasil)
}