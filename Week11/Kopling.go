package main

import "fmt"

func main() {
	var gigi int
	var cop, gas bool

	fmt.Scan(&gigi, &cop, &gas)
	if ((gigi == 0) || (cop && !gas)){
		fmt.Println("Mesin menyala dan motor tidak berjalan")
	} else if (gas){
		fmt.Println("Motor berjalan")
	} else {
		fmt.Println("Mesin mati")
	}
}