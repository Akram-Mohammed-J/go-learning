package main

import "fmt"

// named return values: x and y are declared in the function signature
// a "naked" return returns the named return values without explicitly listing them
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return — returns x and y implicitly
}

// another example with named return values
func rectangle(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // naked return
}

func main() {
	fmt.Println("--- Named Return Values & Naked Return ---")

	x, y := split(17)
	fmt.Println("split(17):", x, y)

	area, perimeter := rectangle(5.0, 3.0)
	fmt.Printf("Rectangle 5x3 → area: %.2f, perimeter: %.2f\n", area, perimeter)
}