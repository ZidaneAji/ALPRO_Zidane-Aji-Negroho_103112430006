package main

import "fmt"

func main(){
	var berat, tinggi, hasil, kilogram float64
	fmt.Print("Masukan Berat Badan : ")
	fmt.Scan(&berat)
	fmt.Print("Masukan Tinggi Badan : ")
	fmt.Scan(&kilogram)

	tinggi = kilogram * kilogram
	berat = berat / berat
	hasil = berat / tinggi

	fmt.Printf("Hasil Nilai BMI : %.2f", hasil)
}