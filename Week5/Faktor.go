package main

import "fmt"

func main() {
	var i, x  int
	var faktor bool

	fmt.Scan(&x)
	for i = 1; i <= x; i++ {
		faktor = x % i == 0
		fmt.Println(i, faktor)
	}
}

// _________________________________________
// Program Faktor
// Kamus
// 	i, x : integer
// 	faktor : boolean
// Algoritma
// 	input(x)
// 	for i <- 1 to x do
// 		faktor <- x mod i == 0
// 		output(i, faktor)
// 	endfor
// Endprogram
// _________________________________________