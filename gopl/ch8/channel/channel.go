package main

import "fmt"

var channel = make(chan int)
var level1 = make(chan int, 100)
var level2 = make(chan int, 100)

func main() {
	//go send(1)
	//x := <-channel
	//fmt.Println("revc:", x)
	go send1(100)
	go recv1()
	go recv1()
	recv2(100)
}

func send(i int) {
	channel <- i
	fmt.Println("send ")
}

func send1(i int) {
	for j := 0; j <= i; j++ {
		level1 <- j
	}
	close(level1)
}

func recv1() {
	for i := range level1 {
		level2 <- i * i
	}
	//close(level2)
}

func recv2(n int) {
	count := 0
	for x := range level2 {
		count++
		fmt.Println(x)
		if n == count {
			close(level2)
		}
	}

}
