package main

import "fmt"

func main() {
    var a, b, hasil int
    
    fmt.Scan(&a, &b)
    
    hasil = 1
    for i := 0; i < b; i++ {
        hasil = hasil * a
    }
    
    fmt.Println(hasil)
}