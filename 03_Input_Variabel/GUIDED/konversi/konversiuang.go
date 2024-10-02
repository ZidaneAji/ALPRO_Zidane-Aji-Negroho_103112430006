package main

import (
	"fmt"
)

func main() {
	var konversi float64 = 15000.0

	var rupiah int

	fmt.Print("Masukkan jumlah uang dalam bentuk Rupiah: ")
	fmt.Scanln(&rupiah)

	dollar := float64(rupiah) / konversi

	fmt.Println("Jumlah uang dalam bentuk Dollar adalah:", dollar)
}