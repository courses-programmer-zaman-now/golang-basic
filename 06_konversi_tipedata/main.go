package main

import "fmt"

func main() {
	// konversi tipe data dari tipe data ke tipe data yang lain

	var nilai32 int32 = 200
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(nilai8)
}
