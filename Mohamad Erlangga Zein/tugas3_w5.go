package main
import (
	"fmt"
	"math"
)

func main(){
	var i, n, n1, n2 float64
	fmt.Print("Masukkan perulangan: ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++{
		fmt.Print("Masukkan bilangan pertama: ")
		fmt.Scan(&n1)
		fmt.Print("Masukkan bilangan kedua: ")
		fmt.Scan(&n2)

		pangkat := math.Pow(n1, n2)

		fmt.Print("Hasil perpangkatannya adalah: ", pangkat)
	}
}