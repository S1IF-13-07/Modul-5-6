package main

import (
	"fmt"
	"math"
)

func main() {
	var radius, height, count int
	fmt.Print("Masukkan jari - jari = ")
	fmt.Scanln(&radius)
	fmt.Print("Masukkan tinggi = ")
	fmt.Scanln(&height)
	fmt.Print("masukkan banyak kerucut = ")
	fmt.Scanln(&count)
	for i := 1; i <= count; i++ {

		luas := 1.0 / 3.0 * math.Pi * float64(radius*radius*height)

		fmt.Printf("Luas kerucut dengan jari - jari %d dan tinggi %d adalah %f\n", radius, height, luas)
	}
}
