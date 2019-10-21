package main

import "fmt"

func main_this2(){
	ans1, ans2, ans3 := multiple_arith(23, 0.04)
	fmt.Println(ans1)
	fmt.Println(ans2)
	fmt.Println(ans3)
}

func evaluate1(){
	a := 24
	var b int = 22

	if (a < b){
		fmt.Println("The first is small")
	} else {
		fmt.Println("The last is small")
	}
}

func addition(a int, b int) int {
	c := a + b
	return c
}

func multiple_arith(a int, b float64) (float64, float64, float64){
	a2 := float64(a) 
	var c float64 = float64(a) * b
	d := a2 * b * a2 
	e := (a2 + b) * a2

	return c, d, e
}