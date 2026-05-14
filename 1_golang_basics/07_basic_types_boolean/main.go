package main 

import ("fmt")

func main() {
	isLogged := true
	isAdmin := false
	hasSubscription := true
	fmt.Println(isLogged) //. inferred as boolean type

	canOpenDashboard  := isAdmin && isLogged

	canDeletePost := isAdmin || (isLogged && hasSubscription)

	fmt.Println(isLogged,isAdmin, hasSubscription )
	fmt.Println(canOpenDashboard, canDeletePost  )
	age := 25

	isAdult := age > 20

	fmt.Println(isAdult)
}