package main

import "fmt"


func message(){
	fmt.Println("Hello World! Go! on the Move!")
}

func variables(){
	var a int
	var b, c float64

	const YEAR int = 1999
	const MONTH int = 12

	a = 23
	b,c = 0.24, 45.44

	fmt.Println(a, " : ", b, " : " , c)

	d := 23444
	fmt.Println(d * 2)
}
