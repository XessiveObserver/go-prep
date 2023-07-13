// Displays smallest and largest number in an array
package main

import "fmt"

func main() {
	var arrayItems, i int
	fmt.Print("Enter an array size to find the smallest and largest=")
	fmt.Scanln(&arrayItems)

	numbArray := make([]int, arrayItems)

	fmt.Println("Enter the array Items=")
	for i = 0; i < arrayItems; i++ {
		fmt.Scanln(&numbArray[i])
	}

	largest := numbArray[0]
	smallest := numbArray[0]

	for i = 0; i < arrayItems; i++ {
		if largest < numbArray[i] {
			largest = numbArray[i]
		}
		if smallest > numbArray[i] {
			smallest = numbArray[i]
		}
	}

	fmt.Println("The Largest Number is", largest)
	fmt.Println("The Smallest Number is", smallest)
}
