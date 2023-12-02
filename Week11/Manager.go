package main

import "fmt"

func main() {
	var m1, m2, m3, m4, m5 string
	fmt.Scan(&m1, &m2, &m3, &m4, &m5)

	if (m1 == "kalah" && m2 == "kalah" && m3 == "kalah" && m4 == "kalah" && m5 == "kalah"){
		fmt.Println("Dipecat")
	} else {
		fmt.Println("Tidak dipecat")
	}
}