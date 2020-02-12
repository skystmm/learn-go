package main

import "fmt"

const(
	black = iota
	red
	_ 
	blue
	green =4
	organge,pink =iota ,iota 
)

func main(){
	fmt.Println("black value: ",black)
	fmt.Println("red value: ",red)
	fmt.Println("blue value: ",blue)
	fmt.Println("green value: ",green)
	fmt.Println("organge value: ",organge)
	fmt.Println("pink value: ",pink)
}