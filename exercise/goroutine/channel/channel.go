package main

import "fmt"

var ch = make(chan int, 1)

func send(a int) {
	ch <- a
}

func recv() {
	i := <-ch
	fmt.Println(i)
}

func main() {
	go send(1)
	recv()
}
