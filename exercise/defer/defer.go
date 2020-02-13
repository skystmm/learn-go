package main

import . "fmt"


func main(){
	Println("first") 
    defer Println("six") 
	Println("second") 
    defer Println("five") 
	Println("third") 
	defer Println("four") 
}