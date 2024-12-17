package main

import "fmt"

func main() {
	var angka int

	for selesai := false; !selesai; {
		fmt.Print("masukan angka : ")
		fmt.Scan(&angka)

		selesai = (angka > 0)
	}
	fmt.Print(angka, " adalah bilangan bulat positif")
}
