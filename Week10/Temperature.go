package main

import "fmt"

func main() {
	var x1, x2, x3, x4, x5 float64

	fmt.Scan(&x1, &x2, &x3, &x4, &x5)
	if ( x1 > x2 && x2 > x3 && x3 > x4 && x4 > x5){
		fmt.Println("Stabil turun")
	} else if ( x1 < x2 && x2 < x3 && x3 < x4 && x4 < x5){
		fmt.Println("Stabil naik")
	} else {
		fmt.Println("Tidak stabil")
	}
}