package main

import "fmt"

func main() {
	var x , y, z int
	var irisan bool

	fmt.Scan(&x, &y, &z)
	irisan = x + z > y
	fmt.Println(irisan)
}