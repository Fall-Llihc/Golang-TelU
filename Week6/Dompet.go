package main

import "fmt"

func main() {
	var saldo, x int

	x = 1
	for x != 0{
		fmt.Scan(&x)
		saldo += x
	}
	fmt.Println(saldo)
}