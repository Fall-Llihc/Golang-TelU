package main

import "fmt"

func main() {
	var i, x, y int

	fmt.Scan(&x)
	for i = 1; i <= x; i++ {
		fmt.Scan(&y)
		fmt.Println(y*y, 4*y)
	}
}

// _________________________________________
// Program Persegi
// Kamus
// 	i, x, y : integer
// Algoritma
// 	input(x)
// 	for i <- 1 to x do
// 		input(y)
// 		output(y*y, 4*y)
// 	endfor
// Endprogram
// _________________________________________