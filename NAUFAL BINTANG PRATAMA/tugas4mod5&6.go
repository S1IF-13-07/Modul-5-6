package main

import "fmt"

func main() {
    var n int
    var i int
    var faktorial int

    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&n)

    faktorial = 1
    i = 1

    for i <= n {
        faktorial = faktorial * i
        i = i + 1
    }

    fmt.Println("Faktorial dari", n, "adalah:", faktorial)
}