package main

import "fmt"

// User is a simple struct with two fields
type User struct {
	name string
	age  int
}

// greet uses a value receiver (u User).
// It receives a COPY of the struct — safe for reading fields.
func (u User) greet() {
	fmt.Println("Hi, I'm", u.name, "and I'm", u.age)
}

// tryBirthday also uses a value receiver.
// Any changes inside this method only affect the copy,
// the original struct in main() stays the same.
func (u User) tryBirthday() {
	u.age++ // only modifies the local copy
	fmt.Println("  Inside method: age =", u.age)
}

func main() {
	// Create a User value
	u := User{name: "Akram", age: 26}

	// Call method — reads from the copy (same values as original)
	u.greet()

	// Call method that tries to modify the struct
	fmt.Println("\nCalling tryBirthday (value receiver)...")
	u.tryBirthday()

	// Original is unchanged because value receiver works on a copy
	fmt.Println("  Original unchanged: age =", u.age)
}