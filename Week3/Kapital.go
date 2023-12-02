package main

import "fmt"

func main() {
	var x int32
	var result bool
	fmt.Scanf("%c", &x)
	
	result = (x >= 65 && x <=90)
	fmt.Println(result)
}