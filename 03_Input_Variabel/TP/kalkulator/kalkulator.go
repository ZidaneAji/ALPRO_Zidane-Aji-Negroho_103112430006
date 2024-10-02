package main

import (
	"fmt"
)

func main() {

	var input float32
	var angkapertama float32
	var angkakedua float32

	fmt.Println("Pilihlah Aritmatika berikut:")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Pembagian")
	fmt.Println("4. Perkalian")
	fmt.Scanln(&input)

	if input == 1 {
		fmt.Println("Masukan angka pertama")
		fmt.Scanln(&angkapertama)
		fmt.Println("Masukan angka kedua")
		fmt.Scanln(&angkakedua)
		fmt.Println("Hasil Penjumlahan adalah", angkapertama+angkakedua)
	} else if input == 2 {
		fmt.Println("Masukan angka pertama")
		fmt.Scanln(&angkapertama)
		fmt.Println("Masukan angka kedua")
		fmt.Scanln(&angkakedua)
		fmt.Println("Hasil Pengurangan adalah", angkapertama-angkakedua)
	} else if input == 3 {
		fmt.Println("Masukan angka pertama")
		fmt.Scanln(&angkapertama)
		fmt.Println("Masukan angka kedua")
		fmt.Scanln(&angkakedua)
		fmt.Println("Hasil pembagian adalah", angkapertama/angkakedua)
	} else if input == 4 {
		fmt.Println("Masukan angka pertama")
		fmt.Scanln(&angkapertama)
		fmt.Println("Masukan angka kedua")
		fmt.Scanln(&angkakedua)
		fmt.Println("Hasil Perkalian adalah", angkapertama*angkakedua)
	} else {
		fmt.Print("Ga ada Pilihan nya ini")
	}
}