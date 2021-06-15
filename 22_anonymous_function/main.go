package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are Blocked ", name)
	}else{
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func (name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("root", blacklist)

	registerUser("danil", func(name string) bool{
		return name == "admin"
	})
}
