package main

import "fmt"

func main() {
	// membuat nama alias tipe data
	type NoKTP string
	type Married bool
	type rumus int32

	var noKtpDanil NoKTP = "32121312112323321"
	var statusNikah Married = true
	var luas rumus = 10 * 3

	fmt.Println(noKtpDanil)
	fmt.Println(statusNikah)
	fmt.Println(luas)
}
