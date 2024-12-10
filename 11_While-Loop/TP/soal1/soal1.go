package main

import "fmt"

func main() {
	var input string
	var maksimum int

	password := "2134"
	maksimum = 3

	for i := 0; i < maksimum; {
		fmt.Print("Masukkan Password yang benar: ")
		fmt.Scan(&input)
		if input == password {
			fmt.Print("Anda Berhasil Masuk")
			return
		} else {
			fmt.Print("Password Anda Salah\n")
			i++
		}

		if i == maksimum {
			print("Anda Melebihi Limit")
		}
	}
}
