package main

import "fmt"

func main(){
	var n int
	fmt.Println("masukan perulangan : ")
	fmt.Scan(&n)

	for i:= 1; i <= n; i++{
		var alas float32
		var tinggi float32
		fmt.Println("masukan tinggi :")
		fmt.Scan(&tinggi)
		fmt.Println("masukan alas : ")
		fmt.Scan(&alas)

		rumus := 0.5 * alas * tinggi

		fmt.Println(rumus)
	}
}