package main

import "fmt"

func main() {
	var masa, anak int
	var jab string

	fmt.Scan(&jab, &masa, &anak)
	if(anak > 3){
		anak = 3
	}
	if (jab == "Staf") {
		if masa < 5 {
			fmt.Println("4000 + 0 = 4000")
		} else if masa > 10 {
			fmt.Printf("5000 + %d = %d", anak*100, 5000+anak*100)
		} else {
			fmt.Printf("4000 + %d = %d", anak*100, 4000+anak*100)
		}
	} else if (jab == "Manager"){
		if masa > 10 {
			fmt.Printf("10000 + %d = %d", anak*300, 10000+anak*300)
		} else {
			fmt.Printf("8500 + %d = %d", anak*300, 8500+anak*300)
		}
	} else {
		fmt.Printf("20000 + %d = %d", anak*500, 20000+anak*500)
	}
}