package main

import "fmt"

func main() {
	var  x, max, min int

	fmt.Scan(&x)
	min = x
	max = x
	
	for i:= 1; i <= 3; i++{
		fmt.Scan(&x)
		if x >= max {
			max = x
		} else if x < min {
			min = x
		}
	}
	fmt.Println(max, min)
}