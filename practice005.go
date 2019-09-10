package main

import "fmt"
//&& and - both sides must be true
//|| or - either side must be true
//! not - is condition false?
func main(){
	fmt.Println(true && true) //True
	fmt.Println(true && false) //False
	fmt.Println(true || true) //True
	fmt.Println(true || false) //True
	fmt.Println(!true) //False
}