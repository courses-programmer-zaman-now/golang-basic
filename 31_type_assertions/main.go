package main

import "fmt"

func random() interface{} {
	return "ini string"
}

func main() {
	var result interface{} = random()

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	case bool:
		fmt.Println("Value", value, "is bool")
	default:
		fmt.Println("Unknown Type")
	}
}
