package main
import "fmt"
func main() {
	var n int
	fmt.Print("Masukan bilangan bulat positif n : ")
	fmt.Scan(&n)

	hasil := 0
	for i := 1; i <= n; i++{
		hasil += i
	}
	fmt.Println("Hasil Penjumlahan : ", hasil)
}