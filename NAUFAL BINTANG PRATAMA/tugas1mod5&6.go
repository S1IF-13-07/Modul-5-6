package main

import "fmt"

func main() {
    var n int
    var i int
    var jumlah int
    fmt.Print("Masukkan nilai n: ")
    fmt.Scan(&n)

    jumlah = 0
    i = 1
    for i <= n {
        jumlah = jumlah + i
        i = i + 1
    }
    fmt.Println("Hasil penjumlahan dari 1 sampai", n, "adalah:", jumlah)
}