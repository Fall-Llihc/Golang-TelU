package main

import "fmt"

func main() {
	var x, d1, d2, d3 int //Comment
	fmt.Scan(&x)
	d1 = x/100
	x = x%100
	d2 = x/10
	x = x%10
	d3 = x
	fmt.Println(d1, d2, d3)
} 