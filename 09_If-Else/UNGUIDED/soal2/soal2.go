/*
Jawaban Soal
a. Keluaran atau Output dari program mengenai nilai 80.1 sebagai inputan maka berpredikat A, sudah sesuai karena pada spesifikasi NAM apabila nilai lebih dari 80 akan menampilkan predikat A
b. pertama, seharusnya pada pemberian kategori untuk predikat itu menggunakan "nmk" dan bukan "nam", seharusnya setelah menggunakan if sebaiknya dilanjutkan dengan penggunaan else if
c. Code dibawah ini adalah code setelah saya perbaiki
*/

package main

import "fmt"

func main() {
	var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)
	if nam > 80 {
		nmk = "A"
	} else if nam <= 80 && nam > 72.5 {
		nmk = "AB"
	} else if nam <= 72.5 && nam > 65 {
		nmk = "B"
	} else if nam <= 65 && nam > 57.5 {
		nmk = "BC"
	} else if nam <= 57.5 && nam > 50 {
		nmk = "C"
	} else if nam <= 50 && nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}
	fmt.Println("Nilai mata kuliah: ", nmk)
}
