package main

import "fmt"

func main() {
	var bil_1, n int
	fmt.Print("Masukan Suatu Bilangan : ")
	fmt.Scan(&bil_1)

	n = 1
	for i := 1; i <= bil_1; i++ {
		n *= i
	}
	fmt.Println("Hasil faktorial dari", bil_1, "adalah", n)
}
