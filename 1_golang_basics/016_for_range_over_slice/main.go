package main

import ("fmt")

func main() {
	views :=[]int {20, 30, 40, 50}
	fmt.Println(views)
	total :=0
	//for range

	for idx, val := range views {
		fmt.Println("day", idx, "views", val)
		total += val
	}

	fmt.Println(total)


}