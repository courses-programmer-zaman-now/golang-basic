package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")
	var username *string = flag.String("user", "admin", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database password")
	var port *int = flag.Int("port", 8080, "Put your database port")

	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("Port : ", *port)
	fmt.Println("User : ", *username)
	fmt.Println("Password : ", *password)

}

// cara memanggil  : go run main.go -host=local -user=danil -password=root
