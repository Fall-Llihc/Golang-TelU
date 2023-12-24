package main

import "fmt"

func main() {
	// var x, i, bin int       			"DENGAN PERKALIAN & TANPA IF"

	// fmt.Scan(&x)

	// i = 1
	// for(x > 0){
	// 	bin = bin + (x%2 * i)
	// 	i *= 10
	// 	x /= 2
	// }
	// fmt.Println(bin) 				"TINGGA DI CONVERT KE STRING"

	var x int
	var y string

	y = ""								//LANGSUNG JADI STRING

	fmt.Scan(&x)
	for x > 0 {
		if x%2 == 0 {
			y = "0" + y
		} else {
			y = "1" + y
		}
		x /= 2
	}
	fmt.Println(y)
}