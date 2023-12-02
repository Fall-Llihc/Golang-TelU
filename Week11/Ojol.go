package main

import "fmt"

func main() {
	var  jam, menit int
	var jarak float64

	fmt.Scanf("%d %d %v", &jam, &menit, &jarak)
	menit = jam*60 + menit

	if (menit >= 300 && menit <= 1320 && jarak >= 0 && jarak <= 20) {
		if((menit > 540 && menit < 960) || (menit > 1140)) {
			fmt.Println(4000*jarak)
		} else if (jarak > 10){
			fmt.Println(4500*jarak)
		} else {
			fmt.Println(5000*jarak)
		}
	} else{
		fmt.Println("Maaf, kami tidak melayani pesanan anda.")
	}
} 