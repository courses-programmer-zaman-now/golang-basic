package main

import "fmt"

func main() {
	// single variable
	var firstName string
	var lastName = "Syah"
	age := 27

	// deklarasi multivariable
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	firstName = "Danil"
	fmt.Println(firstName, lastName, age)

	// multiple variabel
	var (
		namaDepan = "Danil"
		namaBelakang = "Syah"
		umur = 27
	)

	// variabel underscore : variabel untuk menampung nilai yang tidak dipakai
	name, _ := "john", "wick"

	fmt.Println(namaDepan, namaBelakang, "Umur : ", umur)
	fmt.Printf("halo nama saya %s %s dan umur saya %d", namaDepan, namaBelakang, umur)
	fmt.Println(one)
	fmt.Println(isFriday)
	fmt.Println(twoPointTwo)
	fmt.Println(say)
	fmt.Println(name)
}
