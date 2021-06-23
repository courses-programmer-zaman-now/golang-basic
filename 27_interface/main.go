package main

import "fmt"

type HasName interface {
	GetName() string
	GetAge() int
}

func SayHello(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}

func StatusAge(hasname HasName) {
	if hasname.GetAge() >= 17 {
		fmt.Println("Dewasa")
	} else {
		fmt.Println("Kecil")
	}
}

type Person struct {
	Name string
	Age  int
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetAge() int {
	return person.Age
}

type Animal struct {
	Name string
	Age  int
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (animal Animal) GetAge() int {
	return animal.Age
}

func main() {
	var danil Person
	danil.Name = "Danil Syah"
	danil.Age = 27
	SayHello(danil)
	StatusAge(danil)

	dog := Animal{
		Name: "fitbull",
	}
	SayHello(dog)
}
