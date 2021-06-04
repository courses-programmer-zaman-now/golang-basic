package main

import "fmt"

func main() {
	// tipe data numerik non-desimal
	// uint = bilangan cacah (positif)
	// int = bilangan bulat (positif negatif)

	// tipe data numerik desimal = bilangan pecahan
	// float32 dan float64

	var positiveNumber uint8 = 255
	var negativeNumber int8 = -128
	nilai := 999999999999999999

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
	fmt.Println(nilai)

	var ipk = 3.5
	fmt.Printf("nilai IPK saya adalah %f \n", ipk)
	fmt.Printf("nilai IPK saya adalah %.2f \n", ipk)
	// nilai di bulatkan
	fmt.Printf("nilai IPK saya adalah %.f \n", ipk)
}
