
package main

import "fmt"

var x string = "Hello, World!" //global scope, can be used in other functions

func main(){
	fmt.Println(x)
}

func f(){
	fmt.Println(x)
}