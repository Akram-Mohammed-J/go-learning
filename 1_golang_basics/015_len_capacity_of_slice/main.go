package main
import ("fmt")

func main(){
	// slice have two main properties
	//len refers to how many elements you currently have 
	//capacity refers to how many elements you can store

	//make function to create slice 
	//make([]T, LEN, CAP)
	scores := make([]int, 0,5);
	fmt.Println(scores, len(scores), cap(scores))

	scores =append(scores, 100)
	scores = append(scores, 3000)
	scores =append(scores, 300)
	scores = append(scores, 400)
	scores =append(scores, 40)
	fmt.Println("after appending 5 elements", scores, len(scores), cap(scores))
	scores = append(scores, 30)

	//In case of exceeding the capacity, go grows the backing array (usually doubles)
	fmt.Println("go grows the capacity",scores, len(scores), cap(scores))
	scores =append(scores, 100)
	scores = append(scores, 3000)
	fmt.Println("final scores values", scores, len(scores), cap(scores))

	todos := []string{"Make coffee", "learn leetcode", "practice problems", "workout"}
	
	more := []string{"learn go lang"}
	todos = append(todos, more...)
	fmt.Println("Final todos", todos);
}