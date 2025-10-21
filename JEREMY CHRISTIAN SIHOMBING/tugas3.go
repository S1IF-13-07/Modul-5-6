package main
import "fmt"

func main() {
	var a, b int
	var hasil = 1
	fmt.Print("Masukkan dua bilangan: ")
	fmt.Scan(&a, &b)
	for i := 1; i <= b; i++ {
		hasil *= a
	}
	fmt.Println("Hasil:", hasil)
}
