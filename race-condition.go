package main

import (
	"fmt"
	"time"
)

/*
Explanation:
We couldn't specify the order of execution of these 2 methods.
As they are both executed as a GO routine. So, this program
sometimes will print 0 sometimes will print 1. There is a go
command to test this behaviour 'go run -race race-condition.go'
*/

var x int

func PrintX() {
	fmt.Printf("Print X:%d\n", x)
}

func IncrementX() {
	x = x + 1
}

func main() {
	go PrintX()
	go IncrementX()
	time.Sleep(1 * time.Second)

}
