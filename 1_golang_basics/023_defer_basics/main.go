package main
import (
	"fmt"
	"errors"
)


func main() {
// defer -> schedules a function call to run AFTER the surrounding function returns
// even if the function exits early (e.g. due to an error), deferred calls still execute
// common use: cleanup tasks like closing files, releasing locks, or freeing resources

result := doWork(false)
fmt.Println(result)
}

func doWork(success bool) error {
	fmt.Println("Opening a file in simulation..")

	// this will run when doWork returns, regardless of where it returns from
	defer fmt.Println("Cleaning up the state in memory")

	if !success {
		// even though we return early here, the deferred call above still executes
		return errors.New("Something went wrong")
	}
	fmt.Println("Processing the file data")
	fmt.Println("FIle processing Complete")
	return nil
}