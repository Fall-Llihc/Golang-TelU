package main

import "fmt"

func main() {
	var i, x, y int
	fmt.Scan(&x, &y)

	for true {
		i++
		if i%x == 0 && i%y == 0 {
			fmt.Print(i)
			break
		}
	}
}