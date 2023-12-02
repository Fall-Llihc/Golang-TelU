package main

import "fmt"

func main() {
	var user, pass string
	var log int

	for (pass != "admin") && (pass != "admin") {
		fmt.Scan(&user, &pass)
		log += 1
	}
	
	fmt.Println(log-1)
	fmt.Println("Login berhasil")
}