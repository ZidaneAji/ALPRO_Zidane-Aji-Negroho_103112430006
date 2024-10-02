package main

import "fmt"

func main(){
	var sisi int16
	fmt.Print("Masukan Sisi Nya :")
	fmt.Scan(&sisi)

	hasil := sisi * sisi * sisi
	fmt.Printf("volume nya adalah : %d\n", hasil)

}