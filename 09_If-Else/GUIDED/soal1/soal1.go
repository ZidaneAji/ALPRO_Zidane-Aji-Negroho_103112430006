package main

import (
	"fmt"
)

func main() {
	var umur int
	var kartu bool

	fmt.Print("Masukkan usia: ")
	fmt.Scan(&umur)
	fmt.Print("Apakah memiliki kartu keluarga (true/false): ")
	fmt.Scan(&kartu)

	if umur >= 17 && kartu {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}
}
