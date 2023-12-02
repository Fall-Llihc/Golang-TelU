package main

import "fmt"

func main() {
	var x, y, FPB int
	fmt.Scan(&x, &y)

	FPB = 1
	for i := 1; i <= x; i++ {
		if x%i == 0 && y%i == 0 {
			if i > FPB {
				FPB = i
			}
		}
	}
	fmt.Println(FPB)
}