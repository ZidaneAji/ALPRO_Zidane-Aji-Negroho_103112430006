package main

import "fmt"

func main() {
	var namaTanaman string

	fmt.Scan(&namaTanaman)

	switch namaTanaman {
	case "nepenthes", "drosera":
		fmt.Println("Tanaman Karnivora.")
		fmt.Println("Asli Indonesia.")
	case "venus", "sarracenia":
		fmt.Println("Tanaman Karnivora.")
		fmt.Println("Bukan Asli Indonesia.")
	default:
		fmt.Println("Tidak termasuk Tanaman Karnivora.")
	}
}
