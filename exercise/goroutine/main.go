package main

import (
	"fmt"
	_ "learn-go/exercise/goroutine/happensbefore"
)

func init() {
	fmt.Println("init goroutine main package")
}

func main() {
	// import happens before验证
	fmt.Print("start")
}
