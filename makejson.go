package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var inputMap = make(map[string]string)
	fmt.Println("please enter a name:")
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println(err)
	} else {
		inputMap["name"] = name
		fmt.Println("please enter an address:")
		var address string
		_, err := fmt.Scan(&address)
		if err != nil {
			fmt.Println(err)
		} else {
			inputMap["address"] = address
			marshal, err := json.Marshal(inputMap)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(marshal))
			}
		}
	}
}
