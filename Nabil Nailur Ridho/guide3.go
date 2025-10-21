package main
import "fmt"
func main(){
	var j, v1, v2 int
    var hasil int
	fmt.Print("Masukan angka pertama: ")
	fmt.Scan(&v1)
	fmt.Print("Masukan angka kedua: ")
    fmt.Scan(&v2)
    hasil = 0
    for j = 1; j <= v2; j+=1 {
    hasil = hasil + v1
 }
    fmt.Println(hasil)
}