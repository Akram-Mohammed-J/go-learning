package main
import ("fmt")

// with single return value
func  add(a int, b int) int {
	return a + b
}

// with multiple return values

func sumAndProduct(a int, b int) (int,  int) {
	sum := a + b
	product := a * b

	return sum, product
	
}

func main() {
	result := add(20, 30)
	fmt.Println(result)

	//caller of multiple return functions

	sum, product := sumAndProduct(10, 30)
	fmt.Println("Sum and product of 10, 30 is ", sum, product)

	sum2, _ := sumAndProduct(30, 30)
	// use same function but just igonore one return 
	fmt.Println("Sum of 30, 30 is ", sum2)


}