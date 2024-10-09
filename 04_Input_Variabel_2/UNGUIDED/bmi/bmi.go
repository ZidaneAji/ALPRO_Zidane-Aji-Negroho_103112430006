package main

import "fmt"

func main() {
	var bmi, tinggi float64

	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scanln(&bmi)

	fmt.Print("Masukkan tinggi badan (meter): ")
	fmt.Scanln(&tinggi)

	beratBadan := bmi * tinggi * tinggi

	fmt.Printf("Berat badan Anda adalah: %.0f kg\n", beratBadan)
}