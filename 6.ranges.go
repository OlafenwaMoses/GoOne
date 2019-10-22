package main

import (
	"fmt"
)

func main(){
	result := map_from_range()
	fmt.Println(result)
}

func arrays_from_range() []int {
	/* Mapping array values from {} */
	array1 := []int{23, 54, 11, 1, 2, 34}

	return array1
}

func map_from_range() map[string]int{
	/* Mapping map values from {} */
	map1 := map[string]int{"a": 10, "b" : 20, "c" : 30, "d" : 40}

	return map1
}