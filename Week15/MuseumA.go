package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, pls, sum, ttl int
	var mean float64

	fmt.Scan(&x)
	for x >= 0 && x <= 200 {
		ttl++
		sum += x
		fmt.Scan(&y)
		if y > x && y >=0 && y <= 200 {
			pls++
		}
		x = y
		mean = float64(sum) / float64(ttl)
	}
	mean = math.Round(mean*100)/100.0
	fmt.Println(pls, mean)
}