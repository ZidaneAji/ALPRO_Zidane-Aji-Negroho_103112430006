package main

import (
	"fmt"
)

func main() {
	var usia int
	fmt.Print("Masukkan usia Anda: ")
	fmt.Scanln(&usia)

	switch {
	case usia >= 0 && usia <= 12:
		fmt.Println("Kategori usia: Anak-anak")
	case usia >= 13 && usia <= 17:
		fmt.Println("Kategori usia: Remaja")
	case usia >= 18 && usia <= 64:
		fmt.Println("Kategori usia: Dewasa")
	case usia >= 65:
		fmt.Println("Kategori usia: Lansia")
	default:
		fmt.Println("Usia tidak valid")
	}
}
