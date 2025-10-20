package main
import (
	"fmt"
	"math"
)

func main(){
	var n int
	var jariJari, tinggi float64
	fmt.Print("Masukkan perulangan sebanyak n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i += 1{
		fmt.Print("Masukkan nilai jari jari: ")
		fmt.Scan(&jariJari)
		fmt.Print("Masukkan nilai tinggi: ")
		fmt.Scan(&tinggi)
		volumeKerucut := (1.0/3.0) * math.Pi * jariJari * jariJari * tinggi
		fmt.Println("Volume kerucut:", volumeKerucut)
	}
}