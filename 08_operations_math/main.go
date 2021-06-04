package main

import "fmt"

func main() {
	// operasi matematika
	a := 9
	b := 2

	c := a + b
	fmt.Println(c)

	c = a - b
	fmt.Println(c)

	c = a * b
	fmt.Println(c)

	var bagi float32 =  float32(a) / float32(b)
	fmt.Println(bagi)

	c = a % b
	fmt.Println(c)
}
