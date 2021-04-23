package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BubbleSort(slice []int) {
	for i := len(slice); i > 0; i-- {
		for j := 1; j < i; j++ {
			if slice[j-1] > slice[j] {
				Swap(slice, j)
			}
		}
	}
}

func Swap(slice []int, index int) {
	tmp := slice[index]
	slice[index] = slice[index-1]
	slice[index-1] = tmp
}

func main() {
	numberSlice := make([]int, 0)
	message := "Please enter upto 10 integers in CSV(Comma-Separated-Format), e.g. 1,14,64,452,66,72,234"
	for {
		fmt.Println(message)
		var input string
		_, err := fmt.Scan(&input)
		if err == nil {
			split := strings.Split(input, ",")
			if len(split) > 10 {
				fmt.Println("The number of numbers cannot be bigger than 10!")
				continue
			} else {
				for _, value := range split {
					number, err := strconv.Atoi(value)
					if err == nil {
						numberSlice = append(numberSlice, number)
					}
				}
				BubbleSort(numberSlice)
				fmt.Printf("Slice In Order: %v\n", numberSlice)
				break
			}
		}
	}
}
