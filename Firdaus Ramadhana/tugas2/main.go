package main 

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah kerucut yg akan dihitung: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var r, t float64
		fmt.Print("Masukkan panjang alas kerucut dan tinggi kerucut dari kerucut ke-",i , ": ")
		fmt.Scan(&r, &t)
		volume := (1.0 / 3.0) * math.Pi * r * r * t
		fmt.Println("Kerucut ke-",i , "Volumenya adalah: ", volume)
	}
}