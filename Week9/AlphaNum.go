package main

import "fmt"

func main() {
	var k byte
	var alphanum bool

	fmt.Scanf("%c", &k)
	fmt.Println(k)
	alphanum = (k>= 48 && k <= 57) || (k >= 65 && k <= 90) || (k >= 97 && k <= 122)
	fmt.Println(alphanum)
}