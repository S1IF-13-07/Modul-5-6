package main

import "fmt"

  func main() {
  var n int
  fmt.Print("Masukkan bilangan bulat positif n: ")
  fmt.Scan(&n)

    jumlah := 0
    for i := 1; i <= n; i++ {
    jumlah += i
  }
    fmt.Println("Hasil penjumlahan dari 1 sampai", n, "adalah", jumlah)
}
