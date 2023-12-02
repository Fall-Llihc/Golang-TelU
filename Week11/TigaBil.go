package main

import "fmt"

func main() {
	var  a, b, c, temp int
	fmt.Scan(&a, &b, &c)
	if a > c {
		temp = a
		a = c
		c = temp
	}
	if b > c {
		temp = b
		b = c
		c = temp
	}
	if a > b {
		temp = a
		a = b
		b = temp
	}
	fmt.Println(c)
}
