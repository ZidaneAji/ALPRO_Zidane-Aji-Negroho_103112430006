/*kamus
    T, V, VolumeSaatIni : integer
algoritma
    input(T)
    VolumeSaatIni = 0

    while VolumeSaatIni < T do
        input(V)
        VolumeSaatIni = VolumeSaatIni + V
        if VolumeSaatIni lebih dari sama dengan T maka
            output("true")
            break
        else
            output("false")
        endif
    endwhile
endprogram*/

package main

import "fmt"

func main() {
	var T int
	var V int
	fmt.Scan(&T)

	VolumeSaatIni := 0

	for VolumeSaatIni < T {

		fmt.Scan(&V)

		VolumeSaatIni += V

		if VolumeSaatIni >= T {
			fmt.Println("true")
			break
		} else {
			fmt.Println("false")
		}
	}
}
