package main

import "fmt"

func main(){
	var totalharga, diskon, hasil int64

	fmt.Print("masukan total harga : ")
	fmt.Scan(&totalharga)
	fmt.Print("masukan diskon : ")
	fmt.Scan(&diskon)

	hasil = totalharga-(totalharga*diskon/100)
	fmt.Print("Harga Setelah Diskon : ", hasil)

	
}