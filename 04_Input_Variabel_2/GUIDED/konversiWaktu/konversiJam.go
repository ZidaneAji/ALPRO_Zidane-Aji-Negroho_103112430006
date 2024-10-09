package main

import  "fmt"

func main(){
	var (jam,menit,detik,masukan int64)
	
	fmt.Println("Masukan Detik : ")
	fmt.Scan(&masukan)

	jam = masukan / 3600
	menit = masukan / 60
	detik = masukan

	fmt.Println("Hasil Konversi :" , jam, "Jam", menit, "Menit", detik, "Detik")
}