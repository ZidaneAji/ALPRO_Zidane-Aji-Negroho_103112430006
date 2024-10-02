package main

import "fmt"

func main(){
	var alas int
	var tinggi int
	fmt.Print("Masukan Alas Nya :")
	fmt.Scan(&alas)
	fmt.Print("Masukan Tinggi Nya :")
	fmt.Scan(&tinggi)

	hasil := 0.5 * float64(alas) * float64(tinggi)
	fmt.Printf("Luasnya adalah: %.2f\n", hasil)

}