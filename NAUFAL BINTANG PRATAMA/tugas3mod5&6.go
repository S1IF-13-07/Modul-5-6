package main

import "fmt"

func main() {
    var a int
    var b int
    var i int
    var hasil int

    fmt.Print("Masukkan bilangan dan pangkatnya (a b): ")
    fmt.Scan(&a, &b)

    hasil = 1
    i = 1

    for i <= b {
        hasil = hasil * a
        i = i + 1
    }

    fmt.Println("Hasil dari", a, "pangkat", b, "adalah:", hasil)
}