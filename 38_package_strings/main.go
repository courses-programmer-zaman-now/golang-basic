package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Danil Syah", "Danil"))
	fmt.Println(strings.Contains("Danil Syah", "Haykal"))

	fmt.Println(strings.ToLower("Danil Syah Arihardjo"))
	fmt.Println(strings.Split("Danil Syah Arihardjo", " "))
	fmt.Println(strings.ToUpper("haykal dafiansyah"))
	fmt.Println(strings.ToTitle("nufika fitriani"))

	fmt.Println(strings.Trim("=	Danil syah	=", " "))
	fmt.Println(strings.ReplaceAll("danil danil haykal danil", "danil", "fika"))
}
