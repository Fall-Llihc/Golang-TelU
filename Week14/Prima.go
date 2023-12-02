package main

import "fmt"

func main() {
	var x int
	var prime bool

	prime = true
	fmt.Scan(&x)
	if(x>1){
		for i := 2; i <= x/2; i++ {
			if(x%i == 0){
				prime = false
				break
			}
		}
		fmt.Println(prime)
	} else{
		prime = false
		fmt.Println(prime)
	}
}