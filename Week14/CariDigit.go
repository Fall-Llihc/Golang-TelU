package main

import "fmt"

func main() {
	var x, y int
	var ada bool

	fmt.Scan(&y, &x)
	for x > 0 {
		if y == x%10 {
			ada = true
			break
		}
		x /= 10
	}
	fmt.Println(ada)
}