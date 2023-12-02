package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		if i == 1 || i == x {
			for j := 1; j <= x; j++ {
				fmt.Print(i, " ")
			}
		} else{
			for j := 1; j <= x; j++ {
				if j==1 || j ==x{
					fmt.Print(i, " ")
				} else {
					fmt.Print("  ")
				}
			}
		}
		fmt.Println()
	}
}