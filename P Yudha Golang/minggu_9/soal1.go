/*kamus
    mahasiswa, mobil, remaining : integer
algoritma
    kapasitas = 7
    input(mahasiswa)
    mobil = mahasiswa div kapasitas
    remaining = mahasiswa mod kapasitas
    if remaining > 0 maka
        mobil = mobil + 1
        output(mobil - 1, " mobil penuh dan 1 mobil berisi ", remaining, " orang")
    else
        output(mobil, " mobil penuh")
    endif
endprogram*/

package main

import (
	"fmt"
)

func main() {

	const kapasitas = 7

	var mahasiswa int
	fmt.Println("Masukkan jumlah mahasiswa:")
	fmt.Scanf("%d", &mahasiswa)

	mobil := mahasiswa / kapasitas
	remaining := mahasiswa % kapasitas

	if remaining > 0 {
		mobil++
		fmt.Printf("%d mobil penuh dan 1 mobil berisi %d orang\n", mobil-1, remaining)
	} else {
		fmt.Printf("%d mobil penuh\n", mobil)
	}
}
