package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter a string!")
	var input string
	fmt.Scan(&input)
	inputLower := strings.ToLower(input)
	if strings.Contains(inputLower, "i") && strings.Contains(inputLower, "a") && strings.Contains(inputLower, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
