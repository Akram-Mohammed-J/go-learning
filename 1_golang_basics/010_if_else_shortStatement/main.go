package main
import "fmt"

func main() {


	items := 3
	pricePerItem  := 45

	if total := items * pricePerItem ; total > 100 {
		fmt.Println("You are eligible for free shipping")
	} else {
		fmt.Println("you are not eligible for shipping")
	}
}