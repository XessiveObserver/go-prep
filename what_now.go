// functions

package main

import (
	"fmt"
	"math"
	"strings"
)

// struct
type Student struct {
	studentName string
	age         int
	village     string
}

// function
func personality(age int, name string, height float64) {
	fmt.Println(strings.Title(name), "of", age, "years of age\nis", math.Ceil(height), "meters high")
}

// medthod implementing struct Student
func (s Student) id() {
	fmt.Println(strings.Title(s.studentName), "of", s.age, "years of age\nis from", strings.Title(s.village))
}

func profile() {
	personality(22, "okello", 5.8) // call for personality function
	student := Student{"opoo", 25, "omodo"}
	student.id() // call for id method
}
func main() {
	profile()
}
