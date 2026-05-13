package main

import ("fmt")

func main () {
	amount1 := 1000
	amount2 := 2000;

	likes := 10
	likes++
	likes++
	total := amount1+amount2
	avgAmount := total / 2

	fmt.Println("Amount", total)
	fmt.Println(likes, total, avgAmount)

	rating1 := 4.5
	rating2 := 5.1

	avgRating := (rating1 +rating2) / 2

	fmt.Println(rating1, rating2, avgRating)


}