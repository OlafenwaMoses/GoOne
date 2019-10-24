package main

import (
	"fmt"
)

func main() {

	var num int
	num = 26
	mutate(&num)

	
}

func mutate(value *int) {
	*value += 1
}
