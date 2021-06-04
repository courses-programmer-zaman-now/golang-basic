package main

import "fmt"

func main() {
	// constant adalah : sebuah variabel yang nilai nya tidak bisa di ubah / tetap
	// ketika variable constant sudah di deklarasikan nilannya, maka nilainya tidak dapat di ubah

	// mutiple constant
	const (
		firstName = "danil"
		lastName  = "syah"
		age       = 27
	)

	const tinggi = 165
	const nama = "danil"

	fmt.Println(tinggi)
	fmt.Println(nama)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
}
