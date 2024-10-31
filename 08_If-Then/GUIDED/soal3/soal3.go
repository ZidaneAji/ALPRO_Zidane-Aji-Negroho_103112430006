package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukkan angka anda: ")
	fmt.Scan(&angka)

	if angka < 0 && angka%2 == 0 {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}