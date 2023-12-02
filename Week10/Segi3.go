package main

import "fmt"

func main() {
	var b1, b2, b3 int

	fmt.Scan(&b1, &b2, &b3)
	if (b1 == b2 && b2 == b3){
		fmt.Println("Segitiga Sama Sisi")
	} else if (b1 == b2 || b2 == b3 || b3 == b1) {
		fmt.Println("Segitiga Sama Kaki")
	} else {
		fmt.Println("Segitiga Sembarang")
	}
}