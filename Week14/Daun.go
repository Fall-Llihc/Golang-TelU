package main

import "fmt"

func main() {
	var x, y, max int
	fmt.Scan(&x)
	for i := 1; i <= x; i++ { 
		fmt.Scan(&y)
		if y > max {
			max = y
		}
	}
	fmt.Println(max)
}

//---------------- Using Minimal -----------------------
// package main

// import "fmt"

// func main() {
// 	var x, y, max, min int
// 	fmt.Scan(&x)
// 	for i := 1; i <= x; i++ { 
// 		fmt.Scan(&y)
// 		if y > max {
// 			max = y
// 		}
// 		if (min == 0 || y < min){
// 			min = y
// 		}
// 	}
// 	fmt.Println(min)
// 	fmt.Println(max)
// }