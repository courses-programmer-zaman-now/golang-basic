package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 75

	// var lulusUjian = ujian >= 80
	// var lulusAbsensi = absensi >= 80
	// fmt.Println(lulusUjian)
	// fmt.Println(lulusAbsensi)

	// var lulus = lulusAbsensi && lulusUjian
	// var lulus2 = lulusAbsensi || lulusUjian
	fmt.Println(ujian >= 80 && absensi >= 80)
	fmt.Println(ujian >= 80 || absensi >= 80)
}
