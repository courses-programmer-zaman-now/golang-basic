package main

import "fmt"

func main() {
	name := "Danil"
	switch name {
	case "Danil":
		fmt.Println("Hello", name)
	case "Udin":
		fmt.Println("Hello", name)
	default:
		fmt.Println("Kenalan dong")
		fmt.Println("Sini kita salaman")
	}

	// short statement
	switch length := len(name); length > 5 {
		case true:
			fmt.Println("Nama Terlalu panjang")
		case false:
			fmt.Println("Nama sudah benar")
	}

	// tanpa kondisi
	var nilai int = 10
	switch {
		case nilai > 50:
			fmt.Println("Anda Lulus")
		case nilai <= 50 :
			fmt.Println("Anda tidak lulus")
		default :
			fmt.Println("Perbaikan")
	}
}
