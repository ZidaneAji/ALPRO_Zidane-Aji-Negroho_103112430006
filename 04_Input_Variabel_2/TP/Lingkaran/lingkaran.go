package main

import "fmt"

func main() {
	var jari float64
	const pi = 3.14159

	fmt.Print("Masukkan jari-jari : ")
	fmt.Scan(&jari)

	luas := pi * jari * jari

	keliling := 2 * pi * jari

	fmt.Println("Luas :", luas)
	fmt.Println("Keliling : ", keliling)
}