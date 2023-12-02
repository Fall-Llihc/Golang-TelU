package main

import "fmt"

func main() {
	var x, y int
	var kons bool

	kons = true
	fmt.Scan(&x)
	for x >= 10 {
		y = x / 10
		kons = kons && ((x%10 - y%10 == 1) || (x%10 - y%10 == -1))
		x /= 10
	}
	fmt.Println(kons)
}