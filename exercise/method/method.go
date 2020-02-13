package main

import "fmt"

type Human struct {
	name string
	sex  bool
	age  int
}

func (human Human) CheckSex() string {
	if human.sex == true {
		return "男"
	} else {
		return "女"
	}
}

func main() {
	boy := Human{"Tom", true, 30}
	girl := Human{"Luck", false, 18}

	fmt.Println("Tom sex :", boy.CheckSex())
	fmt.Println("Luck sex :", girl.CheckSex())
}
