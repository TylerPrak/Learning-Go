package main

import "fmt"
//Define multiple variables
var (
	a = 1
	b = 2
	c = 2
)

func main(){
	fmt.Println("a =", a) // Seems like Go adds a space after the '=' for me
}