package main

import "fmt"

func main() {
	var a, b, c int
	var x, y int8
	
	fmt.Scanf("%c%c", &x, &y)
	for ((x == 65 || x == 66 || x == 67) && y == 32) {
		if (x == 65) {
			a += 1
		} else if (x == 66) {
			b += 1
		} else if x == 67 {
			c += 1 
		}
		fmt.Scanf("%c%c", &x, &y)
	}
	fmt.Printf("Tipe A = %d\nTipe B = %d\nTipe C = %d", a, b, c)
}