package main

import "fmt"

func main() {
	var x, y, sum int
	var penuh bool

	fmt.Scan(&x)
	for penuh == false {
		fmt.Scan(&y)
		sum += y
		penuh = sum >= x
		fmt.Println(penuh)
	}
}