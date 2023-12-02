package main

import "fmt"

func main() {
	var i, x, n1, n2, nth int
	
	fmt.Scan(&x)
	n1 = 0
	n2 = 1

	for i=0; i<=x; i++ {
		fmt.Print(n1, " ")
		nth = n1 + n2
		n1 = n2
		n2 = nth
	}
}