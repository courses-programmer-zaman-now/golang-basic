package main

import (
	"fmt"
)


func Ups(i int, bebas interface{}) interface{} {
	if i == 1 {
		return 1
	} else if i == 2  {

		return bebas
	} else {
		return "Ups"
	}
}

func main() {
	var data interface{} = Ups(2, "test")
	fmt.Println(data)
}
