package main

import "fmt"

func main() {
	var x int
	var y int
	var hasil int

	fmt.Println("masukan bilangan bulat X :")
	fmt.Scan(&x)
	fmt.Println("masukan bilangan bulat Y :")
	fmt.Scan(&y)

	for i := x; i <= y; i++ {
		hasil = x + y
		fmt.Println("jumlah : ", hasil)
	}

}
