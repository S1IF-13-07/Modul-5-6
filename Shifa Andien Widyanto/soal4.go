package main
import "fmt"
func main() {
	var n int
	fmt.Print("Masukan bilangan bulat : ")
	fmt.Scan(&n)
	
	faktor := 1
	for i :=1; i <= n; i++{
		faktor *= i
	}
	fmt.Println("Hasil faktorialnya adalah : ", faktor)
}
