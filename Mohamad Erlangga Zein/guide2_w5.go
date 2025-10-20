package main
import "fmt"

func main() {
	var j, alas, tinggi, n int
	var luas float64
	fmt.Print("Masukkan perulangan sebanyak n: ")
	fmt.Scan(&n)
	
	for j = 1; j <= n; j += 1{
		fmt.Print("Masukkan panjang sisi alas: ")
		fmt.Scan(&alas)
		fmt.Print("Masukkan panjang sisi tinggi: ")
		fmt.Scan(&tinggi)
		luas = 0.5 * float64(alas * tinggi)
		fmt.Println("luas segitiganya adalah:", luas)
	}
}