package main

import "fmt"

func main(){
	var x,y string
	x = "Hello, "
	y = "World"
	fmt.Println(x==y) //Compares if equal values - false
	
	y = x
	fmt.Println(x==y) //true since y=x

}