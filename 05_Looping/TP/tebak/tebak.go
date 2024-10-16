package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// agar angka acak berubah setiap kali program dijalankan
	rand.Seed(time.Now().UnixNano())

	// Program memilih angka acak antara 1 hingga 100
	target := rand.Intn(100) + 1
	var tebak int

	// Memberikan 5 kesempatan untuk menebak
	for attempts := 1; attempts <= 5; attempts++ {
		fmt.Printf("Tebakan %d: Masukkan angka antara 1 hingga 100: ", attempts)
		fmt.Scan(&tebak)

		if tebak < target {
			fmt.Println("Terlalu kecil!")
		} else if tebak > target {
			fmt.Println("Terlalu besar!")
		} else {
			fmt.Println("Selamat! Tebakanmu benar!")
			break
		}

		// jika masih salah saat percobaan telah berakhir
		if attempts == 5 {
			fmt.Println("Kamu telah mencoba 5 kali. Angka yang benar adalah :", target)
		}
	}
}
