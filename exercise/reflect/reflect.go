package main

import (
	"fmt"
	"reflect"
)

func main() {
	stringType := reflect.TypeOf(1)
	fmt.Println(stringType.Kind())
	fmt.Println(stringType.Name())
}
