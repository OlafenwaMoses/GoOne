package main

import "fmt"

func main_this(){
	/*
	result_array, sample := sample_array()
	fmt.Println(result_array)
	fmt.Println(result_array[3])
	fmt.Println("=========")
	fmt.Println(sample)

	for a := 0; a < len(result_array); a++ {
		fmt.Println(result_array[a])
	}
	*/

	result := multi_dimensional()
	fmt.Println(result)
}

func sample_array() ([5]int, [5]string){
	var array1 [5]int
	array1[3] = 34
	sample := [5]string{"Me", "You", "Us", "We", "Ours"}
	return array1, sample
}

func multi_dimensional() ([2][10][3] int){
	var array [2][10][3]int
	for i := 0; i < 2; i++{
		for j := 0; j < 10; j++{
			for k := 0; k < 3; k++{
				array[i][j][k] = i * j * k
			}
		}
	}

	return array

}