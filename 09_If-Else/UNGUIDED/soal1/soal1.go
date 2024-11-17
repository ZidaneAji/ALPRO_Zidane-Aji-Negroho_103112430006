package main

import "fmt"

func main() {
	var beratG int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratG)

	beratKG := beratG / 1000
	sisaGram := beratG % 1000

	biayaKG := beratKG * 10000
	sisaBiaya := 0

	if beratKG <= 10 {
		if sisaGram >= 500 {
			sisaBiaya = sisaGram * 5
		} else {
			sisaBiaya = sisaGram * 15
		}
	}

	fmt.Printf("Detail Berat: %d KG + %d GR\n", beratKG, sisaGram)
	fmt.Printf("Detail Biaya: Rp. %d + Rp. %d\n", biayaKG, sisaBiaya)
	fmt.Printf("Total Biaya: Rp. %d\n", biayaKG+sisaBiaya)
}
