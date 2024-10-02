package main

import (
	"fmt"
)

func main() {

	var r float32
	var luas float32
	var volum float32

	fmt.Println("Masukan jari-jari bola")
	fmt.Scanln(&r)
	luas = 4 * 3.14 * r * r          //rumus untuk mencari luas bola
	volum = 3.14 * r * r * r / 3 * 4 // rumus untek mencari volumen bola
	fmt.Println("Bola dengan jari-jari", r, "memiliki volum", volum, "dan luas kulit", luas)
}