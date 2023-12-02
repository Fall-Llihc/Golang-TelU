package main

import "fmt"

func main() {
	var n, m, x, y, kopi int

	fmt.Scan(&n, &m, &x, &y)
	for (n > 0) && (m > 0){
		kopi += 1
		n -= x
		m -= y
	}
	fmt.Print(kopi)
}