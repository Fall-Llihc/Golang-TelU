package main

import "fmt"

func main() {
	var x, y, z int
	var segitiga bool

	fmt.Scan(&x, &y, &z)
	segitiga = (x+y >= z) && (y+z >= x) && (x+z >= y)
	fmt.Println(segitiga)
}