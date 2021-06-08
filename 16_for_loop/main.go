package main

import "fmt"

func main() {
	// for
	counter := 1

	for counter <= 10 {
		fmt.Println("nilai ke ", counter)
		counter++
	}

	// for statement
	for index := 1; index <= 10; index++{
		fmt.Println("perulangan ke-",index)
	}

	// for range

	// deklarasi slice
	names := []string{"danil","fika","haykal","khalinda"}

	// cara manual
	for i := 0; i < len(names); i++{
		fmt.Println(names[i])
		fmt.Println("------------")
	}

	// for range
	for index, value := range names{ //gunakan underscore _ jika variable tidak ingin di pakai
		fmt.Println("data ke-",index," = ",value)
	}

	person := map[string]string{
		"name":    "Danil",
		"address": "Sumedang",
		"age":     "27",
	}

	for index, value := range person{
		fmt.Println(index, value)
	}
}
