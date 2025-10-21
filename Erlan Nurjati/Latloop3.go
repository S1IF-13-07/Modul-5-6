package main
import "fmt"
func main() {
    var a, b int
	fmt.Print("Masukan bilangan a: ")
    fmt.Scan(&a)
	fmt.Print("Masukan bilangan b: ")
	fmt.Scan(&b)
    	hasil := 1
   			for i := 0; i < b; i++ {
        	hasil *= a 
    }

    fmt.Println(hasil)
}
