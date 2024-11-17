package main

import "fmt"

func main() {
	var angka int64
	var angka1, angka2, angka3, angka4 int64

	fmt.Print("Masukan 4 digit bilangan : ")
	fmt.Scan(&angka)

	if angka >= 1000 && angka <= 9999 {
		angka1 = angka / 1000
		angka2 = (angka / 100) % 10
		angka3 = (angka / 10) % 10
		angka4 = angka % 10

		if angka1 < angka2 && angka2 < angka3 && angka3 < angka4 {
			fmt.Println("Terurut Membesar")
		} else if angka1 > angka2 && angka2 > angka3 && angka3 > angka4 {
			fmt.Println("Terurut Mengecil")
		} else {
			fmt.Println("Tidak Terurut")
		}
	} else {
		fmt.Print("Harus 4 digit")
	}
}
