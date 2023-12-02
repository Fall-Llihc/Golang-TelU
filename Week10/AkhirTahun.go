package main

import "fmt"

func main() {
	var total int
	var kartu, diskon, cashb bool

	fmt.Scan(&total, &kartu)
	diskon = total >= 100000
	kartu = diskon && kartu
	cashb =  (total >= 200000) && kartu

	if diskon {
		total = total - (total/ 10)
	}
	if cashb {
		total -= 75000
	}

	fmt.Println("Kartu?", kartu) 
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashb)
	fmt.Println("Rp.", total)
}

//________________________________________________________________________
// Program jika mengandung koma
// package main

// import "fmt"

// func main() {
// 	var total int
// 	var nominal float64
// 	var kartu, kartu, diskon, cashb bool

// 	fmt.Scan(&total, &sedia)
// 	diskon = total >= 100000
// 	kartu = diskon && sedia
// 	cashb =  (total >= 200000) && kartu

// 	if diskon {
// 		nominal = float64(total) * 0.9
// 		if cashb {
// 			nominal -= 75000
// 		}
// 		fmt.Println("Kartu?", kartu) 
// 		fmt.Println("Diskon?", diskon)
// 		fmt.Println("Cashback?", cashb)
// 		fmt.Printf("%s%.2f", "Rp.", nominal)
// 	} else if cashb {
// 		total -= 75000
// 		fmt.Println("Kartu?", kartu) 
// 		fmt.Println("Diskon?", diskon)
// 		fmt.Println("Cashback?", cashb)
// 		fmt.Println("Rp.", total)
// 	}
// }