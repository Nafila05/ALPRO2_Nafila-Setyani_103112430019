package main

import "fmt"

func isKabisat(year int) bool {
	return (year%400 == 0) || (year%4 == 0 && year%100 != 0)
}

func main() {
	var tahun int
	fmt.Print("Tahun: ")
	fmt.Scanln(&tahun)

	fmt.Println("Kabisat:", isKabisat(tahun))
}
