package main

import "fmt"

func main() {
	var x, biaya int

	fmt.Scan(&x)
	if x / 1000 < 10 {
		fmt.Printf("%d kg + %d gr\n", x/1000, x%1000)
		biaya = (x/1000) * 10000
		x = x % 1000
		if x >= 500 {
			fmt.Printf("Rp. %d + Rp. %d = Rp. %d ", biaya, x*5, biaya + x*5)
		} else {
			fmt.Printf("Rp. %d + Rp. %d = Rp. %d ", biaya, x*15, biaya + x*15)
		}
		
	} else {
		fmt.Printf("%d kg + %d gr\n", x/1000, x%1000)
		biaya = (x / 1000) * 10000
		fmt.Printf("Rp. %d + Rp. 0 = Rp. %d ", biaya, biaya)
	}
}