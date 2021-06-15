package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// struct method / function
func (customer Customer) sayHello(name string){
     fmt.Println("Hi", name, "my name is", customer.Name )
}

func main() {
     // deklarasi nilai struct
	var pelanggan Customer
	pelanggan.Name = "Danil"
	pelanggan.Address = "Tanjung Uban"
	pelanggan.Age = 27

	fmt.Println(pelanggan)
	fmt.Println(pelanggan.Name)
	fmt.Println(pelanggan.Age)
	fmt.Println(pelanggan.Address)

     // deklarasi nilai struct
	pelanggan2 := Customer{
		Name: "Joko",
		Address: "Cirebon",
		Age: 20,
	}
	fmt.Println(pelanggan2)

	pelanggan3 := Customer{"Budi","Jakarta",45}
	fmt.Println(pelanggan3)

     pelanggan.sayHello("udin")
     pelanggan2.sayHello("icih")
     pelanggan3.sayHello("emen")

}
