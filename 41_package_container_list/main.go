package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Danil")
	data.PushBack("Haykal")
	data.PushBack("Fika")
	data.PushFront("Suparman")

	fmt.Println(data.Back().Value)
	fmt.Println(data.Front().Value)
	fmt.Println(data.Front().Next().Value)

	fmt.Println(data.Front().Prev())
	fmt.Println(data.Back().Next())

	fmt.Println("")

	// dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println("")

	// dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

}
