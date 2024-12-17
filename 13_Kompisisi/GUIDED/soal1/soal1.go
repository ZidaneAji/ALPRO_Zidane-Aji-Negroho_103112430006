package main

import (
	"fmt"
)

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan ganjil : ")
	fmt.Scan(&bilangan)

	fmt.Println("Bilangan ganjil dari 1 hingga", bilangan, "adalah:")
	for i := 1; i <= bilangan; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
	}
}
