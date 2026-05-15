package main 
import ("fmt")

func main(){
	// slices are common collection type
	//dynamic and it can grow
	results := []string{"fouzi", "Akram", "Aaisha"}
	fmt.Println(results[0], len(results)-1)
	results = append(results, "Abrar")
	fmt.Println(results)

}