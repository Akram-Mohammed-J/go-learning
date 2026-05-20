package main
import (
	"fmt"
	"strconv"
	"log"
)


func main() {
	// go dont use exceptions for normal failures
	// functions -> return errors as normal return values

	output, err := parseLevel("rvv")
	if err != nil {
		log.Fatal(err)
	} else {

		fmt.Println(output, err)
	}
}


func parseLevel (s string) (int, error) {

	level, err :=  strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("Value must be a number")
	} else {
		return level, nil
	}

}