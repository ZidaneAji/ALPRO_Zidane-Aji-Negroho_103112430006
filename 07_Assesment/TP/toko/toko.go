package main

import "fmt"

func main() {
	var jumlahBarang, totalPoin int

	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlahBarang)

	// Menghitung poin menggunakan looping
	for i := 1; i <= jumlahBarang; i++ {
		if i <= 5 {
			totalPoin += 10
		} else {
			totalPoin += 15
		}
	}

	fmt.Println("Total poin yang didapatkan:", totalPoin, "poin")
}