package main
import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan satu bilangan: ")
	fmt.Scan(&n)
	var hasil = 1 
	for i := 1; i <= n; i++{
		hasil = hasil * i
	}

	fmt.Print("Memiliki Hasil: ", hasil)
}