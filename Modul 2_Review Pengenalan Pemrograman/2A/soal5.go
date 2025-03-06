package main

import "fmt"

func main() {
	var nums [5]int
	var input string

	// Membaca 5 angka integer
	fmt.Print("Masukkan 5 angka ASCII: ")
	for i := 0; i < 5; i++ {
		fmt.Scan(&nums[i])
	}

	// Membaca 3 karakter langsung sebagai string
	fmt.Print("Masukkan 3 karakter tanpa spasi: ")
	fmt.Scan(&input)

	// Konversi dan cetak 5 angka ke karakter
	for i := 0; i < 5; i++ {
		fmt.Printf("%c", nums[i])
	}
	fmt.Println()

	// Cetak 3 karakter berikutnya
	fmt.Println(input)
}
