package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)
	if x % 5 == 0 {
		fmt.Println("Kelipatan 5")
	}
	if x % 3 == 0 {
		fmt.Println("Kelipatan 3")
	}
}