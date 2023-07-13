// Takes users name and age and displays it to console

package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Whast's your name?")
	fmt.Scanln(&name)
	fmt.Println("Whast's your age?")
	fmt.Scanln(&age)

	fmt.Println(name, "is your name", age, "years of age")
}
