package main

import "fmt"

func main() {
	var target, donasi, totalDonasi, jumlahDonatur int

	// Input target donasi
	fmt.Print("Masukkan target donasi: ")
	fmt.Scanln(&target)

	// Mulai dengan loop yang setara dengan repeat-until
	for {
		// Minta input jumlah donasi
		fmt.Print("Masukkan jumlah donasi: ")
		fmt.Scanln(&donasi)

		// Tambahkan donasi ke total
		totalDonasi += donasi
		jumlahDonatur++

		// Cetak status terkini
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, totalDonasi)

		if totalDonasi >= target {
			break
		}
	}

	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonasi, jumlahDonatur)
}
