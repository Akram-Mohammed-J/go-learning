package main

import ("fmt")

func addBonusBy5(score *int) {
	*score +=5
}

func main() {
	//    & -> this points the address of the data stored in the memory
	//   * -> this points to value that stored in a particular addess in the memory

	score:= 10
	fmt.Println("Before adding bonus score value is :", score )
	addBonusBy5(& score)
	fmt.Println("After adding bonus by 5, score value is :", score )



}