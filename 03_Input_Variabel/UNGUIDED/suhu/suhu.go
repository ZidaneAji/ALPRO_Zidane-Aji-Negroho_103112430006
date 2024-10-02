package main

import "fmt"

func main() {
	var celcius float32
	var fahrenheit float32
	var reamur float32
	var kelvin float32

	fmt.Println("masukkan temperatur celcius :")
	fmt.Scanln(&celcius)

	fahrenheit = (celcius * 9.0 / 5.0) + 32
	reamur = (4.0 / 5.0) * celcius
	kelvin = celcius + 273.15

	fmt.Println("Temperatur Celcius :", celcius)
	fmt.Println("Derajat Reamur :", reamur)
	fmt.Println("Derajat Fahrenheit :", fahrenheit)
	fmt.Println("Derajat Kelvin :", kelvin)

}