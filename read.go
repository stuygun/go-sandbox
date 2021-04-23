package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	maxLength = 20
)

type Name struct {
	fname string
	lname string
}

func (n *Name) trim(s string, length int) string {
	var size, x int

	for i := 0; i < length && x < len(s); i++ {
		_, size = utf8.DecodeRuneInString(s[x:])
		x += size
	}
	return s[:x]
}

func (n *Name) Set(firstName string, lastName string) {
	n.fname = n.trim(firstName, maxLength)
	n.lname = n.trim(lastName, maxLength)
}

/*
	Write a program which reads information from a file and represents it in a slice of structs.
	Assume that there is a text file which contains a series of names.
	Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

	Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
	Each field will be a string of size 20 (characters).

	Your program should prompt the user for the name of the text file.
	Your program will successively read each line of the text file and create a struct
	which contains the first and last names found in the file. Each struct created will be added to a slice,
	and after all lines have been read from the file, your program will have a slice containing
	one struct for each line in the file. After reading all lines from the file, your program should iterate through
	your slice of structs and print the first and last names found in each struct.
*/
func main() {
	nameSlice := make([]Name, 0)
	fmt.Println("please enter the name of the text file:")
	var fileName string
	_, e := fmt.Scan(&fileName)
	if e != nil {
		fmt.Println(e)
	} else {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			split := strings.Split(line, " ")
			obj := &Name{}
			obj.Set(split[0], split[1])
			nameSlice = append(nameSlice, *obj)
		}

		for i, v := range nameSlice {
			fmt.Printf("%d. First-Name: %s, Last-Name: %s\n", i+1, v.fname, v.lname)
		}
	}
}
