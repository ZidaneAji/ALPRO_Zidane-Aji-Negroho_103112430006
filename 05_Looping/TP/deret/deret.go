package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	// Variabel untuk menyimpan hasil penjumlahan
	sum := 0

	// Melakukan perulangan dari 1 hingga n dan menjumlahkan
	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println("Jumlah dari 1 hingga", n, "adalah :", sum)
}