package main
import (
	"fmt"
	"strconv"
)

/*
Work like dictionaries
*/

func main(){

	basic_map()

}

func basic_map(){
	sample := make(map[string]int)

	for a:=0; a < 20; a++{
		sample[strconv.Itoa(a) ] = a * 23
	}

	fmt.Println(sample)
}