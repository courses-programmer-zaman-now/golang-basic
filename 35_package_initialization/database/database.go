package database

import "fmt"

var connection string

// secara otomatis akan di jalankan ketika package di panggil
func init() {
	fmt.Println("init dipanggil")
	connection = "MySQL"
}

func GetConnection() string {
	return connection
}
