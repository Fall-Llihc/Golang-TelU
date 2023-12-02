package main

import "fmt"

func main() {
	var umur int
	var KK bool

	fmt.Scan(&umur, &KK)
	if (umur >= 17 && KK) {
		fmt.Println("Bisa membuat KTP")
	} else {
		fmt.Println("Belum bisa membuat KTP")
	}
}