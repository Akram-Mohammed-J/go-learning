package main
import "fmt"

type User struct {
	name string
	age int
}
func main() {
	u := User {
		name: "Akram",
		age: 27,
	}
	fmt.Println("Age before update ..: ", u.age)
	u.updateAge();
	fmt.Println("Age After update ..: ", u.age)
}

func (u *User) updateAge() {
	u.age++
}