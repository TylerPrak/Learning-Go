package main

import "fmt"

func main(){
	var x string = "Hello, World"
	fmt.Println(len(x)) //Reads length of string variable x
	fmt.Println(x[1]) //Grabs second index of string variable x
	fmt.Println("Hello, " + "World") //Concatenates strings
}