package main

import "fmt"

/*
Works like arrays, but with option to append new items
*/

func main(){
	result := slices(6)
	fmt.Println(result)
	result = append(result, 1000)
	fmt.Println(result)
}


func slices(a int) []int{
	slice := make([]int, a)
	for aa := 0; aa < a; aa++{
		slice[aa] = aa * 2
	}
	
	return slice
}
