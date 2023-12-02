package main

import "fmt"

func main() {
	var x int
	var kabisat bool

	fmt.Scan(&x)
	kabisat = (x%400 != 0 && x%100 != 0 && x % 4 == 0) || (x%400 == 0)
	fmt.Println(kabisat)
}