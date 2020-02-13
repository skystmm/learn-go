package main

import "fmt"

type Human struct {
	name string
	sex  bool
	age  int
}

type Employee struct {
	Human
	salary float32
	deaprt string
	age    int
}

func main() {
	employee := Employee{Human: Human{name: "Sky", sex: true}, salary: 100.00}
	employee.age = 30
	employee.Human.age = 30
	employee.deaprt = "go learn"

	fmt.Println(employee)

	employee.name = "Test"
	fmt.Println(employee)
}
