package main

import "fmt"

func main() {

	// cara pertama deklarasi array
	var names [3]string

	names[0] = "Danil"
	names[1] = "Haykal"
	names[2] = "Fika"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// cara kedua deklarasi array
	var values = [3]int{
		90,
		70,
		66,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// mengubah data pada posisi index
	values[0] = 75
	fmt.Println(values)

	// function build in
	fmt.Println("jumlah data : ", len(values))
}
