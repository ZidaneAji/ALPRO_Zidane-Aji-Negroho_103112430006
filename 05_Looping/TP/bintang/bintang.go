package main

import (
	"fmt"
)

func main() {
	var baris int

	// Meminta input
	fmt.Print("Masukkan jumlah baris segitiga bintang: ")
	fmt.Scan(&baris)

	// Mencetak segitiga bintang
	for i := 1; i <= baris; i++ {
		// Mencetak spasi untuk membuat segitiga di tengah
		for j := 1; j <= baris-i; j++ {
			fmt.Print(" ")
		}
		// Mencetak bintang untuk membentuk segitiga simetris
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}