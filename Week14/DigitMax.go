package main

import "fmt"

func main() {
	var x, max int
	fmt.Scan(&x)
	for x > 0{ 
		if x%10 > max {
			max = x%10
		}
		x /= 10
	}
	fmt.Println(max)
}