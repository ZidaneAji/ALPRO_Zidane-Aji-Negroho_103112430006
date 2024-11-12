package main

import (
	"fmt"
	"strings"
)

func main() {
	var huruf string
	fmt.Print("Masukkan huruf: ")
	fmt.Scan(&huruf)

	huruf = strings.ToUpper(huruf)

	if huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" {
		fmt.Println("Vokal")
	} else {
		fmt.Println("Konsonan")
	}
}
