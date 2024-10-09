package main

import "fmt"

func main() {
	var jamMingguan float64
	var upahPerJam float64

	fmt.Print("Masukkan jumlah jam kerja dalam satu minggu: ")
	fmt.Scan(&jamMingguan)

	fmt.Print("Masukkan upah per jam : RP ")
	fmt.Scan(&upahPerJam)

	var jamNormal, jamLembur float64
	if jamMingguan > 40 {
		jamNormal = 40
		jamLembur = jamMingguan - 40
	} else {
		jamNormal = jamMingguan
		jamLembur = 0
	}

	gajiMingguan := (jamNormal * upahPerJam) + (jamLembur * 1.5 * upahPerJam)

	gajiBulanan := gajiMingguan * 4

	fmt.Println("Total gaji bulanan : RP ", int(gajiBulanan))
}