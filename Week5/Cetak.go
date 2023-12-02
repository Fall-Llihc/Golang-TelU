package main

import "fmt"

func main() {
	var x int
	var y string

	fmt.Scan(&x, &y)
	for i := 1; i <= x; i++ {
		fmt.Println(y)
	}
}

// _________________________________________
// Program CetakString
// Kamus
// 	i, x : integer
// 	text : string
// Algoritma
// 	input(x)
// 	input(text)
// 	for i <- 1 to x do
// 		output(text)
// 	endfor
// Endprogram
// _________________________________________