// displays multiplication table of a number

package main

import "fmt"

func main(){
	var num int

	fmt.Println("Enter any number and\n get it's multiplication table upto 10 ")
	fmt.Scanln(&num)
	for i := 1; i < 11; i++{
		product := num * i
		fmt.Println(product)
	}
}
