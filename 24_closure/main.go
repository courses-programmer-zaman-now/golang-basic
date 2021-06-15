package main

import "fmt"

func main() {
	name := "Danil"
	counter := 0

	// fungsi di dalam fungsi
	// closure adalah memampuan sebuah function berinteraksi dengan data-data disekitarnya dlaam scope
	// yang sama,
	// harap gunakan  fitur closure ini dengan bijak saat kita membuat aplikasi
	increment := func() {
		name = "Haykal"

		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
