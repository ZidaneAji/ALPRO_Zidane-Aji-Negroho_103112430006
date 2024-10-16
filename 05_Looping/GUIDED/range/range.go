package main

import "fmt"

func main(){
	var a int16
	var b int16

	fmt.Print("masukan bilangan bulat awal :")
	fmt.Scan(&a)
	fmt.Print("masukan bilangan bulat tujuan :")
	fmt.Scan(&b)

	for i:=a; i <= b; i++{
		fmt.Print(" ", i)
	}
}