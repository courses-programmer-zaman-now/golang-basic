package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int){
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println(result)
}


func endApp(){
	message := recover()
	if message != nil{
		fmt.Println("Error dengan pesan : ", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool){
	defer endApp()
	if error{
		panic("ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}


func main() {
	// runApplication(10)
	runApp(true)
	fmt.Println("program terus berjalan")
}
