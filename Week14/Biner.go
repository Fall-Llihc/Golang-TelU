package main

import "fmt"

func main() {
	var x, i, bin int

	fmt.Scan(&x)

	i = 1
	for(x > 0){
		bin = bin + (x%2 * i)
		i *= 10
		x /= 2
	}
	fmt.Println(bin) 
}