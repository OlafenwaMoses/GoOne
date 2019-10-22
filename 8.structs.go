package main

import (
	"fmt"
)

type worker struct{
	name string
	age int
	salary float64
}

func main(){

	workerX := create_worker("James Bond", 43, 24300.0)

	fmt.Println(workerX.name)
	fmt.Println(workerX.age)
	fmt.Println(workerX.salary)

}

func create_worker(name string, age int, salary float64) worker{
	new_worker := worker{}
	new_worker.name = name
	new_worker.age = age
	new_worker.salary = salary

	return new_worker
}