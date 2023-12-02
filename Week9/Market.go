package main

import "fmt"

func main() {
	var total int
	var sedia, kartu, diskon, cashb bool

	fmt.Scan(&total, &sedia)
	diskon = total >= 100000
	kartu = diskon && sedia
	cashb =  (total >= 200000) && kartu

	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashb)
}