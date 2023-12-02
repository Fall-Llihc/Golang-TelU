package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)
	if ( x < 0){
		x *= -1
	}
	fmt.Println(x)	
}