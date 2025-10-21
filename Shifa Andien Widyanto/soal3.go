package main
import (
	"fmt"
	"math"
)
func main() {
	var x, y float64
	fmt.Print("Masukan dua bilangan positif : ")
	fmt.Scan(&x, &y)
	hasil := math.Pow(x, y)
	fmt.Println("Jadi hasil pemangkatannya adalah : ", hasil)
}