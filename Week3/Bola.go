package main

import "fmt"

func main() {
	var r, L float64
	fmt.Scan(&r)
	L = 4 * 22/7.0 * (r*r)
	fmt.Println(L)
}