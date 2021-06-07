package main

import "fmt"

func main() {
	name := "udin"
	isMarried := false

	if name == "Danil" && isMarried {
		fmt.Println("Hello", name)
	}else if name == "udin" && !isMarried {
		fmt.Println("kamu blm menikah", name)
	}else{
		fmt.Println("nama tidak ditemukan")
	}

	// short statement
	if length := len(name); length > 5{
		fmt.Println("nama terlalu panjang")
	}else{
		fmt.Print("nama sudah benar")
	}
}
