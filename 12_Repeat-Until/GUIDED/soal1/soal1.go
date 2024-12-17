package main

import "fmt"

func main() {
	var word string
	var repetitions int

	fmt.Print("Masukkan kata & banyaknya perulangan : ")
	fmt.Scan(&word, &repetitions)

	for i := 1; i <= repetitions; {
		fmt.Println(word)
		i++
	}
}
