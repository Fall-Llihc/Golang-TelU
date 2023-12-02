package main

import "fmt"

func main() {
	var t, jt, mt, dt int

	fmt.Scan(&t)
	jt = t / 3600
	t %= 3600
	mt = t / 60
	t %= 60
	dt = t

	fmt.Println(jt, "jam", mt, "menit", dt, "detik")
}