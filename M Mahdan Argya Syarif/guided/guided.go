package guided

import "fmt"

func Guided1(a int, b int) {
	var j int
	for j = a; j <= b; j += 1 {
		fmt.Print(j, " ")
	}
}

func Guided2(n int, alas int, tinggi int) {
	var j int
	var luas float64
	for j = 1; j <= n; j += 1 {
		luas = 0.5 * float64(alas*tinggi)
		luas = 0.5 * float64(alas*tinggi)
		fmt.Println(luas)
	}
}
