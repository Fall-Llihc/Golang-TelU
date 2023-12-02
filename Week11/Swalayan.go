package main

import (
	 "fmt"
	 "math"
)

func main() {
	var mem string
	var a, b, c, d , e int
	var cb, ds float64

	fmt.Scan(&mem, &a, &b, &c, &d, &e)
	if (a% 2 != 0 && b% 2 != 0 && c% 2 != 0 && d% 2 != 0 && e% 2 != 0){
		ds = 1.7 * float64(c+d+e)

	} else if (a% 2 == 0 && b% 2 == 0 && c% 2 == 0 && d% 2 == 0 && e% 2 == 0){
		cb = 3.1 * float64(a+b+c)

	} else {
		cb = 3.1 * float64(a+b+c)
		ds = 1.7 * float64(c+d+e)
	}
	if mem == "Yes" {
		ds += ds*0.15
		cb += cb*0.15
	}
	if (cb > 35){
		cb = 35
	}
	if ds > 50 {
		ds = 50
	}
	cb = math.Round(cb*100) /100.0 //Guasah ditulis cuma buat buletin angka
	ds = math.Round(ds*100) /100.0

	fmt.Printf("Cashback: %v%% Diskon: %v%%", cb, ds) //output("Cashback: ", cd, "% Diskon: ", ds, "%")
}