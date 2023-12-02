package main

import (
	"fmt"
)

func main() {
	var x, sum int

	for i := 1; i <= 5; i++ {
		fmt.Printf("Hari ke %d : ", i)
		fmt.Scan(&x)
		for x < 0 || x > 200 {
			fmt.Printf("Hari ke %d : ", i)
			fmt.Scan(&x)
		}
		sum += x
	}
	fmt.Printf("Jumlah pengunjung : %d orang", sum)
}