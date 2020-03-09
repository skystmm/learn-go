package main

import (
	"fmt"
	"learn-go/gopl/ch12/display"
	"reflect"
	"unsafe"
)

type ReflictDemo struct {
	Reflict int
	test    string
	demos   map[Demo]int
	Demo
}

type Demo struct {
	demo     int
	testDemo float32
}

func main() {
	demo1 := Demo{demo: 1, testDemo: 1.3}
	re := ReflictDemo{Reflict: 1, test: "123", Demo: Demo{demo: 123, testDemo: 32.2}}
	demos := make(map[Demo]int)
	demos[demo1] = 123
	re.demos = demos
	rt := reflect.TypeOf(re)
	dt := reflect.TypeOf(re.Demo)
	fmt.Println(rt)
	fmt.Println(dt)
	fmt.Printf("%v\n", reflect.ValueOf(re))
	display.Displayhans("a", demos)

	sizeof := unsafe.Sizeof(re)
	fmt.Print(sizeof)
}
