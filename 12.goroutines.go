package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	speaker(4, "I am ")

	go speaker(4, "You are ")

	go func(msg string) {
		fmt.Println(msg)
	}("WE ARE GOOD")

	time.Sleep(time.Second)
	fmt.Println("DONE")

}

func speaker(times int, msg string) {
	for a := 0; a < times; a++ {
		fmt.Println(msg, strconv.Itoa(a))
	}
}
