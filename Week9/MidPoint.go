package main

import "fmt"

func main() {
	var b1, b2, b3 int
	var c1, c2, c3 bool
	
	fmt.Scan(&b1, &b2, &b3)
	c1 = (b2 + b3) % 2 == 0 && (b2 + b3) / 2 == b1
	c2 = (b1 + b3) % 2 == 0 && (b1 + b3) / 2 == b2
	c3 = (b1 + b2) % 2 == 0 && (b1 + b2) / 2 == b3

	fmt.Println(c1 || c2 || c3)
}