package main

import "fmt"

func main() {
	var i, x, sum int

	fmt.Scan(&x)
	sum = 1
	for i = x; i >= 1; i-- {
		sum = sum * i
	}
	fmt.Println(sum)
}

// _________________________________________
// Program Faktorial
// Kamus
// 	i, x, sum : integer
// Algoritma
// 	input(x)
// 	sum <- 1
// 	for i <- x down to 1 do
// 		sum <- sum * i
// 	endfor
// 	output(sum)
// Endprogram
// _________________________________________
