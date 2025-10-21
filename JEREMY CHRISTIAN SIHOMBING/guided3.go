package main
import "fmt"

func main()	{
	var a, b int
	var hasil = 0
	fmt.Print("Input A: ")
	fmt.Scan(&a)
	fmt.Print("Input B: ")
	fmt.Scan(&b)
	for i := 1; i <= a; i++ {
		hasil += b 
	}
	fmt.Println("Hasil:", hasil)
}
