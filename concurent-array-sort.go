package main

import (
	"fmt"
	"sort"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var arrLen int
	fmt.Println("Enter length of array")
	fmt.Scan(&arrLen)

	arr := make([]int, arrLen)

	var number int
	for i := 0; i < arrLen; i++ {
		fmt.Println("arr: ", arr)
		fmt.Println("Enter number")
		fmt.Scan(&number)
		arr[i] = number
	}

	//arr = []int{2,43,65,1,2579,9,12,5343,65,87, 43,1,5,6,8,2,3,4}

	chunkSize := len(arr) / 4

	if lenMod := len(arr) % 4; lenMod != 0 {
		chunkSize += 1
	}

	if chunkSize < 1 {
		fmt.Println("Error: length of array is less than 4")
	}
	j := 0

	fmt.Println("Before sorting: arr: ", arr, "chunkSize: ", chunkSize, "length: ", len(arr))
	for i := 0; i <= len(arr); i += chunkSize {
		if i >= len(arr) {
			break
		}
		fmt.Println("i : ", i)
		last := i + chunkSize
		if last > len(arr) {
			last = len(arr)
		}
		if last > i {
			wg.Add(1)
			go sortSubArr(arr[i:last], j)
			j += 1
		}
	}

	wg.Wait()
	fmt.Println("subarrays sorted: ", arr)
	sort.Ints(arr)
	fmt.Println("full array sorted: ", arr)
}

func sortSubArr(arr []int, index int) {
	defer wg.Done()
	sort.Ints(arr)
	fmt.Println("GoRoutine#", index, "sorted subArray: ", arr)
}
