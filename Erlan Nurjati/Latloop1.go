package main
import "fmt"
func main() {
    var n int
    var hasil int 
    fmt.Print("Masukkan bilangan n: ")
    fmt.Scan(&n)
	hasil = 0
    for i := 1; i <= n; i++ {
        hasil += i //sama dengan hasil = hasil + i
    }
    fmt.Println(hasil)
}