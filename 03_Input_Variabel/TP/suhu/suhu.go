package main

import (
	"fmt"
)

func main() {

	var farenheit float32
	var rumus float32

	fmt.Println("Masukan suhu farenheit : ")
	fmt.Scanln(&farenheit)
	
	rumus = (farenheit * 5 / 9) + 273
	fmt.Println("Jadi suhu kelvin nya adalah", rumus)
}