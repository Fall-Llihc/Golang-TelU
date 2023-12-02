package main
import "fmt"

func main(){
	var x, y int
    var f float64
    fmt.Scan(&x)
    fmt.Scan(&y)

    f = ((1/float64((3*(x*x)+10)))+float64((10*y))+7)
    fmt.Println(f)
}