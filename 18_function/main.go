package main

import "fmt"

func tambah(a int, b int) int {
	return a + b
}

func bagi(a int, b int) float32{
	return float32(a) / float32(b)
}

func sayHello(){
	fmt.Println("heloo world")
}

// function parameter dengan return
func cekKelulusan(nilai int)string{
	if nilai > 65{
		return "anda lulus"
	}else{
		return "anda gagal"
	}
}

// function dengan parameter
func greeting(name string){
	fmt.Println("hai ", name)
}

// returning multiple values
func getFullName()(string, string, string){
	return "Danil","syah","arihardjo"
}

func penjumlahan(bil1 int, bil2 int)(string, int){
	return "penjumlahan", bil1 + bil2
}

// function return named value
func getFullName2()(firstName string, middleName string, lastName string, age int){
     firstName = "Danil"
     middleName = "Ari"
     lastName = "Hardjo"
     age = 27

     return
}

func main() {
	sayHello()
	result := tambah(10, 2)
	fmt.Println(result)
	fmt.Println(bagi(3,2))
	greeting("danil syah")
	fmt.Println(cekKelulusan(64))

	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)

	title, result := penjumlahan(5,3)
	fmt.Println(title, result)

     a, b, c, d := getFullName2()
     fmt.Printf("Nama %s %s %s umur saya %d ", a, b, c, d)
}
