package main

import "fmt"

// membuat nama alias
type Filter func(string, int) string

// fungsi sebagai parameter
func sayHelloWithFilter(name string, age int, filter Filter ) {
	nameFiltered := filter(name, age)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string, age int) string{
	if name == "Anjing" && age <= 17{
		return "umur dibawah 17 tahun tidak boleh sebut Anjing"
	}else{
		return name
	}
}

func main() {
	sayHelloWithFilter("Anjing",17, spamFilter)
	sayHelloWithFilter("Danil",18 ,spamFilter)
}
