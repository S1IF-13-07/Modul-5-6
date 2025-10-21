package unguided

import "math"

func Unguided1(a int) int {
	var i int
	var total int
	total = 0
	for i = 1; i <= a; i += 1 {
		total += i
	}
	return total
}

func Unguided2(n int, r float64, t float64) float64 {
	var total float64
	for i := 0; i < n; i++ {
		volume := (1.0 / 3.0) * math.Pi * math.Pow(r, 2) * t
		total += volume
	}
	return total
}

func Unguided3(a int, b int) int {
	result := 1
	for i := 1; i <= b; i++ {
		result *= a
	}
	return result
}

func Unguided4(n int) int {
	faktorial := 1
	for i := 1; i <= n; i++ {
		faktorial *= i
	}
	return faktorial
}
