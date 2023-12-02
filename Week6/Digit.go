package main

import "fmt"

func main() {
	var x, sum int
	
	fmt.Scan(&x)
	for x >= 1 {
		fmt.Print(x%10, " ")
		sum += x%10
		x /= 10
	}
	fmt.Println()
	fmt.Println(sum + x)
}