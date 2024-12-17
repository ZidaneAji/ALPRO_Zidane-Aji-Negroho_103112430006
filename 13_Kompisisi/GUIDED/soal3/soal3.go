package main

import (
	"fmt"
)

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&bilangan)

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
