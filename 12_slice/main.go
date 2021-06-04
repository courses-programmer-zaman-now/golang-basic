package main

import (
	"fmt"
)


func main() {
	// deklarasi array
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[2] = "editttt"
	fmt.Println(slice1)
	fmt.Println(months)

	fmt.Println("=====================================================")

	days := [...]string{"senin","selasa","rabu","kamis","jumat","sabtu","minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Libur"
	daySlice1[1] = "Minggu Libur"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// membuat slice baru
	newSlice := make([]string, 2, 5)

	newSlice[0] = "danil"
	newSlice[1] = "haykal"

	fmt.Println(newSlice)

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Prin
	tln(iniSlice)
}

