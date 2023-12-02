package main

import "fmt"

func main() {
	var x, y float64

	fmt.Scan(&x, &y)
	if x < y {
		fmt.Println("Peningkatan sebesar", y-x)
	} else if x > y{
		fmt.Println("Penuruan sebesar", x-y)
	} else {
		fmt.Println("Tetap")
	}
}