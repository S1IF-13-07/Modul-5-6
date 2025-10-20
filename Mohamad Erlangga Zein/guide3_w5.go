package main
import "fmt"

func main() {
	var a, b, hasil int
	fmt.Print("Masukkan bilangan pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan kedua: ")
	fmt.Scan(&b)

	hasil = 0

	for i := 1; i <= b; i+=1{
		hasil = hasil + a
	}
	fmt.Print("hasilnya: ", hasil)
}