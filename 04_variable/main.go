package main

import "fmt"

func main() {
	// single variable
	var firstName string
	var lastName = "Syah"
	age := 27

	firstName = "Danil"
	fmt.Println(firstName, lastName, age)

	// multiple variabel
	var (
		namaDepan = "Danil"
		namaBelakang = "Syah"
		umur = 27
	)

	fmt.Println(namaDepan, namaBelakang, "Umur : ", umur)
}
