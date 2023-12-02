package main

import "fmt"

func main() {
	var k byte
	
	fmt.Scanf("%c", &k)
	if (k == 'A' || k == 'I' || k == 'U' || k == 'E' || k == 'O' || k == 'a' || k == 'i' || k == 'u' || k == 'e' || k == 'o') {
		fmt.Println("Bukan Konsonan")
	} else {
		fmt.Println("Konsonan")
	}
}