package main

import (
	"fmt"
	"programmerzamannow/golang-basic/34_package_import/helper"
)


func main() {
	result := helper.SayHello("Danil Syah")
	fmt.Println(result)
	fmt.Println(helper.FirstName)
}
