/*
kamus

	x, j, bilangan, nX, nZero : integer

algoritma

	input(x)
	for j = 1 to 9 do
	    input(bilangan)
	    if bilangan == x maka
	        nX = nX + 1
	    else
	        nZero = nZero + 1
	    endif
	endfor
	if nX > nZero then
	    output("Modus = ", nX)
	else
	    output("Modus = ", nZero)
	endif

endprogram
*/
package main

import (
	"fmt"
)

func main() {
	var x, bilangan, nX, nZero int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	for j := 1; j <= 9; j++ {
		fmt.Print("Masukkan bilangan ke-", " ", j)
		fmt.Scan(&bilangan)
		if bilangan == x {
			nX++
		} else {
			nZero++
		}
	}

	if nX > nZero {
		fmt.Print("Modus = ", x, "(dengan jumlah ", nX, ")")
	} else {
		fmt.Print("Modus = 0 (dengan jumlah ", nZero, ")")
	}
}
