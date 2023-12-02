package main

import "fmt"

func main() {
	var i, x, y, sum int
	fmt.Scan(&x)

	for i = 1; i <= x; i++ {
		fmt.Scan(&y)
		sum += y
	}
	fmt.Printf("%.3f", (float64(sum)/float64(x)))
}

// _________________________________________
// Program Ngoding
// Kamus
// 	i, x, y, sum : integer
// Algoritma
// 	input(x)
// 	for i <- 1 to x do
// 		input(y)
// 		sum <- sum + y
// 	endfor
// 	output(sum/x)
// Endprogram
// _________________________________________
