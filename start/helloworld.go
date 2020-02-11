package main

import (
	"fmt"
	"strconv"
)

type Test struct {
	name string
	test int
}

func (t Test) String() string {
	return "test { name:" + t.name + ",test:" + strconv.Itoa(t.test) + "}"
}

func main() {
	t := Test{"test", 123}
	fmt.Println(t)

	var a = [6]int{1, 2, 3, 4, 5, 6}
	b := a[0:]
	c := a[0:]

	b[5] = 1
	fmt.Println(c[5])

	b = append(b, 7)
	fmt.Println(cap(b))
	b[1] = 9
	fmt.Println(c[1])
}
