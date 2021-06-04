package main

import "fmt"

func main() {
	// operasi matematika
	a := 9
	b := 2

	fmt.Println("==== Operasi Matematika ====")
	c := a + b
	fmt.Println(c)

	c = a - b
	fmt.Println(c)

	c = a * b
	fmt.Println(c)

	var bagi float32 = float32(a) / float32(b)
	fmt.Println(bagi)

	c = a % b
	fmt.Println(c)

	fmt.Println("==== Augmented Assigments ====")
	// augmented assignments
	a += 10
	fmt.Println(a)

	a -= 2
	fmt.Println(a)

	a *= 3
	fmt.Println(a)

	a /= 2
	fmt.Println(a)

	a %= 3
	fmt.Println(a)

	// unary operators
	fmt.Println("=== Unary Operators ===")
	b++
	fmt.Println(b)
	b--
	fmt.Println(b)
	fmt.Println(-b)
	fmt.Println(+b)

	var isMarried = false
	fmt.Println(!isMarried)


}
