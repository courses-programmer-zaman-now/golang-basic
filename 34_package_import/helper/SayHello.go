package helper

var FirstName string //bisa di import
var fullName string  //tidak bisa di import

func SayHello(name string) string { //bisa di import
	return name
}

func intro(name string) string { //tidak bisa di import
	return name
}
