package main

import "fmt"

func main() {
	var basis, pangkat, hasil int

	fmt.Print("Masukkan bilangan pertama untuk basis: ")
	fmt.Scan(&basis)

	fmt.Print("Masukkan bilangan kedua untuk pangkat: ")
	fmt.Scan(&pangkat)

	hasil = 1

	for i := 0; i < pangkat; i++ {
		hasil *= basis
	}

	fmt.Println("Hasil dari", basis, "pangkat", pangkat, "adalah: ", hasil)
}