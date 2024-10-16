package main

import "fmt"

func main() {
	var n int
	hasil := 1

	fmt.Println("masukkan bilangan positif:")
	fmt.Scan(&n)

	for i := 2; i <= n; i++ {
		hasil *= i
	}
	fmt.Print(hasil)
}