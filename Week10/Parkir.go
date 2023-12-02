package main

import "fmt"

func main() {
	var h1, m1, h2, m2, hh, mm int

	fmt.Scan(&h1, &m1, &h2, &m2)
	m1 = (h1 * 60) + m1
	m2 = (h2 * 60) + m2
	mm = m2 - m1

	if (mm <= 0){
		mm = 12 * 60 + mm
		hh = mm/60
		mm %= 60
	} else {
		hh = mm/60
		mm %= 60
	}

	fmt.Printf("%d jam %d menit", hh, mm)
}