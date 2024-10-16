package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	// menginputkan jumlah kerucut
	fmt.Print("Masukkan jumlah kerucut: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var r, t, volume float64

		// Meminta input jari-jari dan tinggi untuk setiap kerucut
		fmt.Printf("kerucut ke-%d, Masukkan jari-jari dan tinggi kerucut: ", i)
		fmt.Scan(&r, &t)

		// Menghitung volume kerucut tanpa fungsi
		volume = (1.0 / 3.0) * math.Pi * r * r * t

		// Menampilkan hasil volume kerucut
		fmt.Printf("Volume kerucut ke-%d: %.5f\n", i, volume)
	}
}