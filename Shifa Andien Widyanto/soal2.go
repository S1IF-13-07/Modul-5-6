package main
import (
	"fmt"
	"math"
)
func main() {
	var n int
	fmt.Print("Masukan bilangan bulat : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++{
		var r, t float64
		fmt.Print("Masukan jari-jari dan tinggi dari kerucut : ")
		fmt.Scan(&r, &t)
		volume := (1.0 / 3.0)*math.Pi*r*r*t
		fmt.Println(volume)
	}
}