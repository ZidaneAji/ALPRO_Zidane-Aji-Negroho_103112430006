package main

import "fmt"

func main() {
	var a int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&a)

	if a < 1 {
		fmt.Print("Bukan Positif")
	} else {
		fmt.Print("Positif")
	}

}