package main

import "fmt"

func main() {
	var jam12, jam24 int
	var label string

	fmt.Print("Input : ")
	fmt.Scan(&jam24)

	switch {
	case jam24 == 0:
		jam12 = 12
		label = " AM"
	case jam24 < 12:
		jam12 = jam24
		label = " AM"
	case jam24 == 12:
		jam12 = 12
		label = " PM"
	case jam24 > 12 && jam24 < 25:
		jam12 = jam24 - 12
		label = " PM"
	default:
		fmt.Print("Tidak Termasuk Jam")
	}
	fmt.Print(jam12, label)
}
