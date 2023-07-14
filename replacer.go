// Go programme to swap characters
package main

import (
	"fmt"
	"strings"
)

func main() {
	brokenWords := "b##ks #f f##d"

	replacer := strings.NewReplacer("#", "o")

	correctWords := replacer.Replace(brokenWords)

	fmt.Println(correctWords)
}
