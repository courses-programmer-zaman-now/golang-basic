package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 2) //conversi int ke binary
	fmt.Println(value)

	intToString := strconv.Itoa(192)
	fmt.Println(reflect.TypeOf(intToString))

	stringToInt, _ := strconv.Atoi("100")
	fmt.Println(reflect.TypeOf(stringToInt))
}

// base size
// desimal : 10
// oktal : 8
// binary : 2
// heksadesimal : 16
