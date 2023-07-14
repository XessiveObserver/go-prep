// Number guessing game

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen random number between 1 and 100")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Make a guess:")
	input, error := reader.ReadString('\n')
	if error != nil {
		log.Fatal(error)
	}
	input = strings.TrimSpace(input)
	guess, error := strconv.Atoi(input)
	if error != nil {
		log.Fatal(error)
	}

	if guess < target {
		fmt.Println("Oops. Your guess was LOW")
	} else if guess > target {
		fmt.Println("Oops. Your guess was HIGH")
	}
}
