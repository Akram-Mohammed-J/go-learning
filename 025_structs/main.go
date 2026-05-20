package main 

import ("fmt")

// struct -> a collection of fields, similar to a class in other languages but without inheritance
// fields starting with uppercase (Id) are exported (accessible from other packages)
// fields starting with lowercase (name, age, address) are unexported (package-private)
type User  struct {
	Id int
	name string
	age int
	address string
}

func main() {
	// creating a struct instance using named fields (order doesn't matter with named fields)
	u1 := User {
		Id: 1,
		name: "Akram",
		age: 26,
		address: "89 chinna arippu street kottar nagercoil",
	}
	fmt.Println(u1)
	// struct fields are mutable by default

	u1.age = 27
	fmt.Println("After updating the age ", u1.age)

	//its not mandatory to initialize all fields in struct 

	u2 := User {
		name:"fouzi",
	}
	fmt.Println("Partial user u2:", u2)

}