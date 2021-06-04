package main

import "fmt"

func main() {
	var ember1 = "air"
	var ember2 = "api"

	var result bool = ember1 == ember2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
