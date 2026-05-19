package main
import "fmt"

func sumAll(nums ...int) int {
	total := 0
	for _, val :=range nums {
		total+= val
	}
	return total
}

func main() {
	r1 := sumAll(1, 2, 3, 4,5)
	fmt.Println("result of summAll .... sumAll(1, 2, 3, 4,5): ",r1)

	input := []int {10, 20, 30}
	r2 :=  sumAll(input...)
	fmt.Println("result of summAll .... {10, 20, 30}: ",r2)

	//Anonymus function
	res:= func (a int) int {
		return a*2;
	}

	//IIFE in Go
	fmt.Println("Anonymus function result: ",res(10))

	 func (input int)  {
		fmt.Println("IIFE CALLED",  input)
	}(10)
}