package main 
import "fmt" 
 
func main() { 
    var j, alas, tinggi, n int 
    var luas float64 
	fmt.Print("Masukkan jumlah data: ")
    fmt.Scan(&n) 
    for j = 1; j <=n; j+=1 { 
        fmt.Print("Masukkan alas dan tinggi: ")
        fmt.Scan(&alas, &tinggi) 
        luas = 0.5 * float64(alas * tinggi) 
        fmt.Println("Luas segitiga: ", luas) 
    } 
} 