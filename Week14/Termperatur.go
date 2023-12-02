package main

import "fmt"

func main() {
	var i, x, y, sum, min, max int

	fmt.Scan(&x, &y)
	max = x
	min = x
	i = 1
	for x!=0 || y != 0 {
		sum += x
		x = y

		if(x > max) {
			max = x
		} else if (x < min) {
			min = x
		}

		fmt.Scan(&y)
		i += 1
	}
	fmt.Println(max)
	fmt.Println(min)
	fmt.Println(float64(sum)/float64(i))
}