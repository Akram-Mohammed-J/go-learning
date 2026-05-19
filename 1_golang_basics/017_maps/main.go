package main
import ("fmt")

func main(){
	//map syntax -> map[keyType]valueType
	ages := map[string]int{
		"Akram": 27,
		"fouzi":23,
	} 
		fmt.Println(ages)

	// the below code wont work if u want to initalize a empty map
	// var scores map[string]int

	// scores["Akram"] = 20
	// fmt.Println("The Score of AK is ", scores["Akram"])
	
		//the correct way of initializing the empty map
		scores := make(map[string]int)
		fmt.Println("The Score map is empty", scores)
		scores["Akram"] = 200
		scores["fouzi"] = 400
		scores["Abrar"] = 300
		scores["Aaisha"] = 200
		fmt.Println("The Score Ak is ", scores["Akram"])

		// delete  a key in map  
		delete(scores, "Aaisha")
		fmt.Println("score map length", len(scores))

		// if you try to delete a key which doesn't exist in map it wont through any error

		delete(scores, "u100") // no error

		// loop over the key values in map
		for key, val := range scores {
			fmt.Println(key, val)
		}



}