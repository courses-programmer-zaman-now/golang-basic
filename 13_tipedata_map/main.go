package main

import (
	"fmt"
)


func main() {
	person := map[string]string{
		"name":    "Danil",
		"address": "Sumedang",
		"age":     "27",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["name"])

	// menambahkan key dan value baru
	person["status"] = "menikah"
	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(len(person))

	// delete data
	delete(person, "status")

	fmt.Println(person)

	// deklarasi map baru
	matkul_sks := make(map[string]int)
	matkul_sks["pemrograman"] = 2
	matkul_sks["inggris"] = 1
	matkul_sks["agama"] = 2

	fmt.Println(matkul_sks)
	fmt.Println("jumlah data ", len(matkul_sks))
}
