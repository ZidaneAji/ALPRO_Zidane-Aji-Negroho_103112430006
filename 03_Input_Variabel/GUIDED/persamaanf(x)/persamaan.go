package main

import (
	"fmt"
)

func main() {

	var f float32
	var rumus float32

	fmt.Println("Masukan angka x")
	fmt.Scanln(&f)
	rumus = (2/(f+5) + 5)
	fmt.Println("Jadi hasil untuk persamaan nya adalah", rumus)
}