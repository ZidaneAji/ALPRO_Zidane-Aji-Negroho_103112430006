package main

import "fmt"

func main() {
	var input, total int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&input)

	// Menghitung jumlah bilangan
	for i := 1; i <= input; i++ {
		total += i
	}

	// Menampilkan hasil keluaran
	fmt.Println("Keluaran: ", total)
}