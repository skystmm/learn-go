package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	abort := interrupt()
	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func interrupt() <-chan struct{} {
	abort := make(chan struct{})
	go func() {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			content := input.Text()
			if content == "exit" {
				abort <- struct{}{}
			}
		}
	}()
	return abort
}

func launch() {
	fmt.Println("launch")
}
