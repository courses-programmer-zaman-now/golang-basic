package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	// menentukan waktu secara manual
	utc := time.Date(2021, 7, 19, 13, 50, 20, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02" //format layout
	parse, _ := time.Parse(layout, "2021-07-19")
	fmt.Println(parse)

	parse2, _ := time.Parse(time.RFC3339, "2021-10-01T5:04:05Z")
	fmt.Println(parse2)
}
