package main
import "fmt"

func main() {
	var n, hasil int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		hasil += i
	}
	fmt.Println("Hasil penjumlahan dari 1 sampai", n, "adalah", hasil)
}
