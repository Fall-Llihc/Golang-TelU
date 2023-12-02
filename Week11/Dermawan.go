package main

import "fmt"

func main() {
	var sum, x, m1, m2, m3, m4, m5 int

	fmt.Scan(&x)
	if x  == 4 {
		fmt.Scan(&m1, &m2, &m3, &m4)
		sum = m1 + m2 + m3 + m4 + m5
	} else if x == 5 {
		fmt.Scan(&m1, &m2, &m3, &m4, &m5)
		sum = m1 + m2 + m3 + m4 + m5
	}
	fmt.Println(sum)
}	