package main
import (
	"fmt"
	"math"
)

func main() {
	var n int
	var r, t float64
	fmt.Print("Masukkan jumlah kerucut: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print("Masukkan jari-jari dan tinggi: ")
		fmt.Scan(&r, &t)
		volume := (1.0 / 3.0) * math.Pi * math.Pow(r, 2) * t
		fmt.Println("Volume kerucut:", volume)
	}
}
