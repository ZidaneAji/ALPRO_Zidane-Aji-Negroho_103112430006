/*kamus
    a, b, c : integer
algoritma
    input(a, b, c)
    if a lebih kecil sama dengan 0 atau b lebih kecil sama dengan 0 atau c lebih kecil sama dengan 0 maka
        output("Sisi segitiga tidak valid. Sisi harus bilangan positif.")
    else if a sama dengan b dan b sama dengan c maka
        output("Segitiga Sama Sisi")
    else if a sama dengan b atau b sama dengan c atau a sama dengan c maka
        output("Segitiga Sama Kaki")
    else
        output("Segitiga Sembarang")
    endif
endprogram*/

package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scanln(&a, &b, &c)

	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("Sisi segitiga tidak valid. Sisi harus bilangan positif.")
	} else if a == b && b == c {
		fmt.Println("Segitiga Sama Sisi")
	} else if a == b || b == c || a == c {
		fmt.Println("Segitiga Sama Kaki")
	} else {
		fmt.Println("Segitiga Sembarang")
	}
}
