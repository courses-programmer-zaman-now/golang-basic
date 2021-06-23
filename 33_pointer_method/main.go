package main

import "fmt"

// membuat struct
type Man struct {
	Name string
}

// struct method
func (man *Man) Married() {
	man.Name = "Mr " + man.Name
}

func main() {
	person := Man{"danil"}
	person.Married()
	fmt.Println(person.Name)
}
