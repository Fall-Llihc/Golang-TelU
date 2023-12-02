package main

import "fmt"

func main() {
	var i, N int
	var x1, x2, x3, x4, x5, status bool

	fmt.Scan(&N)
	status = true
	for i = 1; i <= N; i++ {
		fmt.Scan(&x1, &x2, &x3, &x4, &x5)
		status = status && x1 && x2 && x3 && x4 && x5	
	}
	fmt.Println(status)
}