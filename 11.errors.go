package main

import (
	"errors"
	"fmt"
)

func main() {

	status, error := compare(30)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(status)
	}

}

func compare(value int) (bool, error) {
	if value < 20 {
		return false, errors.New("Value is too low")
	} else {
		return true, nil
	}
}
