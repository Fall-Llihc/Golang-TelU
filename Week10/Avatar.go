package main

import "fmt"

func main() {
	var x, dewasa, kecil int

	fmt.Scan(&x)
	if x <= 15 {
		dewasa = x / 5
		if x % 5 != 0 {
			dewasa += 1
		}
		fmt.Println("Dewasa", dewasa)
	} else if x > 15 && x <= 25 {
		kecil = (x-15) / 2
		if (x-15) % 2 != 0 {
			kecil += 1
		}
		fmt.Printf("Dewasa 3, kecil %d", kecil)
	}  else{
		fmt.Println("Dewasa 3, kecil 5, dan", x-25, "tak berangkat")
	}
}