package main

import "fmt"

func main() {
	var x int
	var dkn, cb bool

	fmt.Scan(&x)
	dkn = ((x/10)%10) % 2 == 0
	cb = (x%10 != 0) && ((x/1000) + ((x/10)%10)) % (x%10) == 0
	fmt.Println("Diskon?", dkn)
	fmt.Println("Casshback?", cb)
}