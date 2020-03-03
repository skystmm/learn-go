package main

import (
	"fmt"
	"strings"
)

/*
去除string中的空格
*/
func NoneEmptyStr(src string) string {
	result := make([]string, len(src))
	index := 0
	for _, c := range src {
		if c != ' ' {
			result[index] = string(c)
			index++
		}
	}
	return strings.Join(result, "")
}

func main() {
	fmt.Println(NoneEmptyStr(" a c b   d"))
}
