package main
import ("fmt"
"strings"

)


func main() {
	firstName := "Akram"
	lastName := "Mohammed"
	fullName := firstName + " " +  lastName

	fmt.Println(fullName)
	fmt.Println(strings.ToUpper(fullName))

}