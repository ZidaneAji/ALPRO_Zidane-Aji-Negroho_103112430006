package main

import "fmt"

func main() {
	var tahun int
	var kabisat int

	fmt.Println("Masukkan tahun :")
	fmt.Scanln(&tahun)

	kabisat = tahun % 4

	if kabisat == 0 {
		fmt.Println("Tahun:", tahun)
		fmt.Println("Kabisat: true")
	} else {
		fmt.Println("Tahun:", tahun)
		fmt.Println("Kabisat: false")
	}
}