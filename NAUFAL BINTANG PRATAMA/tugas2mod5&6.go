package main

import (
    "fmt"
    "math"
)

func main() {
    var n int
    var i int
    var r, t, volume float64

    fmt.Print("Masukkan jumlah kerucut: ")
    fmt.Scan(&n)

    i = 1
    for i <= n {
        fmt.Print("Masukkan jari-jari dan tinggi kerucut ke-", i, ": ")
        fmt.Scan(&r, &t)

        volume = (1.0 / 3.0) * math.Pi * r * r * t
        fmt.Printf("Volume kerucut ke-%d adalah: %.14f\n", i, volume)

        i++
    }
}