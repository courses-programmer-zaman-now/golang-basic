package main

import (
	"flag"
	"fmt"
)


func main() {
	host := flag.String("host","localhost","Put your database host")
	username := flag.String("username","userroot","Put your database username")
	password := flag.String("password","passwordroot","Put your database password")
	number := flag.Int("number", 94, "using number")

	flag.Parse()

	fmt.Println("Host: ", *host)
	fmt.Println("User: ", *username)
	fmt.Println("Password: ", *password)
	fmt.Println("Number", *number)
}
