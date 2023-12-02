package main

import "fmt"

func main() {
	var poin int
	var gold string = "Gold User"
	var platinum string = "Platinum User"
	var silver string = "Silver User"

	fmt.Scan(&poin)
	if poin < 0 {
		poin = 0
	}
	poin += 50
	if poin >=50 && poin <= 99 {
		fmt.Println(silver)
	} else if poin >= 100 && poin <= 200 {
		fmt.Println(platinum)
	} else {
		fmt.Println(gold)
	}
}