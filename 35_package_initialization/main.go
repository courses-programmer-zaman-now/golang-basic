package main

import (
	// _ "programmerzamannow/golang-basic/35_package_initialization/database" // blank identifier
	"fmt"
	"programmerzamannow/golang-basic/35_package_initialization/database"
)


func main(){
	fmt.Println(database.GetConnection())
}
