package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv" // conversion to and from string representations of basic data types
	"strings"
	"sync"
)

func main() {
	fmt.Println("Hello World")
	pointersTest()
	callFunctionPointer()
	printFFunction()
	differentIntegers()
	typeConversions()
	typeFloatAndComplexNumbers()
	stringsPackage()
	strconvPackage()
	//scan()
	arrays()
	slices()
	maps()
	structs()
	ioutils()
	osFileOps()
	variableAsFunction()
	functionAsArguments()
	anonymousFunction()
	testFunctionReturnsFunction()
	testVariadic()
	testDeferred()
	testInterfaces()
	testWaitGroup()
}

func pointersTest() {
	var x int = 1
	var y int
	var pointer *int
	pointer = &x
	y = *pointer
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("pointer address:", pointer)
	fmt.Println("pointer value:", *pointer)

	pointer2 := new(int)
	*pointer2 = 3
	fmt.Println("pointer2 value:", *pointer2)
}

func callFunctionPointer() {
	var y *int
	y = functionPointer()
	fmt.Println("y in callFunctionPointer:", *y)
}

func functionPointer() *int {
	x := 1
	return &x
}

//Print some words
func printFFunction() {
	x := "Joe"
	fmt.Printf("Hi " + x + "\n")
	fmt.Printf("Hi %s", "serhan")
	fmt.Printf("\n")
}

func differentIntegers() {
	/*different lengths and signs:
	int8, int_16, int32, int64
	uint_8, uint_16, uint32, uint64
	defining only int will leave the compiler to choose one of the int
	*/
	var int_8 int8 = 127 // -128, 127
	fmt.Printf("int8: %d", int_8)
	fmt.Printf("\n")
	var int_16 int16 = 32767 // -32768, 32767
	fmt.Printf("int_16: %d", int_16)
	fmt.Printf("\n")
	var uint_8 uint8 = 255 // 0, 255
	fmt.Printf("uint_8: %d", uint_8)
	fmt.Printf("\n")
	var uint_16 uint16 = 65534 // 0, 65536
	fmt.Printf("uint_16: %d", uint_16)
	fmt.Printf("\n")
}

func typeConversions() {
	var x int32 = 1
	var y int16 = 2
	//x = y will give error
	x = int32(y)
	fmt.Printf("typeConversions:%d", x)
	fmt.Printf("\n")

}

func typeFloatAndComplexNumbers() {
	var a float32 = 1.1234567          //6 digits of precision
	var b float64 = 1.1234567890123456 // 15 digits of precision
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	//complex numbers
	var z complex128 = complex(2, 3)
	fmt.Println("z(complext128):", z)
}

func stringsPackage() {
	str1 := "str1"
	str2 := "str2"
	compare := strings.Compare(str1, str2)
	fmt.Println("compare('str1', 'str2'):", compare)

	str3 := "str3"
	subStr3 := "tr3"
	contains := strings.Contains(str3, subStr3)
	fmt.Println("strings.Contains('str3', 'tr3'):", contains)

	str4 := "str4"
	str4Prefix := "st"
	hasPrefix := strings.HasPrefix(str4, str4Prefix)
	fmt.Println("strings.HasPrefix('str4', 'tr'):", hasPrefix)

	str5 := "str5"
	str5SubStr := "tr"
	index := strings.Index(str5, str5SubStr)
	fmt.Println("strings.Index('str5', 'tr'):", index)
}

func strconvPackage() {
	//Some functions of the strconv package:
	//Atoi: converts string to int
	//Itoa: converts int (base10) to string
	//FormatFloat(f, fmt, prec, bitSize): converts floating point number a string
	//ParseFloat(s, bitSize): converts a string to a floating point number

	x, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Error occurring during 'x, err := strconv.Atoi(\"123\")'")
	} else {
		y := x + 3
		fmt.Println("y: strconv.Atoi(\"123\") + 3:", y)
	}
}

func contants() {
	const x = 1.3
	const (
		y = 4
		z = "Hi"
	)

	//iota: generates a set of related but distinct constants like days of week
	type GRADE int
	const (
		A GRADE = iota
		B
		C
		D
		F
	)
}

func controlFlows() {
	//if-else
	test := false
	if test {
		fmt.Println("test is true")
	} else {
		fmt.Println("test is false")
	}
	//for loops
	//1.
	for i := 0; i < 10; i++ {
		fmt.Printf("Hi")
	}
	//2.
	j := 0
	for j < 10 {
		fmt.Printf("hi")
		j++
	}

	//switch
	x := 0
	switch x {
	case 1:
		fmt.Printf("1")
	case 2:
		fmt.Printf("2")
	default:
		fmt.Printf("0")
	}

	//tagless switch
	y := 0
	switch {
	case y >= 1:
		fmt.Printf("y >= 1")
	case y <= -1:
		fmt.Printf("y <= -1")
	default:
		fmt.Printf("y = 0")
	}

	//break
	k := 0
	for k < 10 {
		k++
		if k == 5 {
			break
		}
		fmt.Printf(" hi")
	}
	//continue
	k = 0
	for k < 10 {
		k++
		if k == 5 {
			continue
		}
		fmt.Printf(" hi")
	}
}

func scan() {
	var appleNum int

	fmt.Printf("Number of apples?")
	n, err := fmt.Scan(&appleNum)
	if err == nil {
		fmt.Println("n: ", n)
		fmt.Println("Number of apple entered: ", appleNum)
	} else {
		fmt.Println(err)
	}
}

func arrays() {
	fmt.Println("TESTING ARRAYS")
	//default-initialised arrays
	var arr [5]int
	fmt.Println(arr[0])
	arr[1] = 3
	fmt.Println(arr[1])

	//pre-initialised array
	abc := [5]int{1, 2, 3, 4, 5}
	fmt.Println(abc[4])

	//partially initialised array

	var def = [5]int{10, 20, 30}
	fmt.Println(def[2])
	fmt.Println(def[3])

	//array literal infers size
	ghi := [...]int{1, 2, 3, 4}
	fmt.Println(ghi[0])

	//looping an array
	for i, v := range ghi {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
}

func slices() {
	fmt.Println("TESTING SLICES")
	/*a "window" on an underlying array
	variable size, upto the whole array
	Every slice:
		POINTER indicates the start of the slice
		LENGTH is the number of elements in the slice
		CAPACITY is maximum number of elements
			-From the start of slice to end of array. As the size can be increased
	*/
	array := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Printf("array:={\"a\", \"b\", \"c\", \"d\", \"e\", \"f\", \"g\"}, length: %d, capacity: %d\n", len(array), cap(array))
	slice1 := array[1:3] //zero based index. Start Inclusive, End Exclusive
	slice2 := array[2:5]
	fmt.Printf("slice1:=array[1:3] length: %d, capacity: %d\n", len(slice1), cap(slice1))
	fmt.Printf("slice2:=array[2:5] length: %d, capacity: %d\n", len(slice2), cap(slice2))
	fmt.Printf("slice1[1]: %s\n", slice1[1])
	fmt.Printf("slice2[0]: %s\n", slice2[0])

	//Difference between array and slice literal initialisation
	// arr := [3] int {1,2,3}
	// slc := [] int {1,2,3} -> compiler understands from empty square bracket length it is a slice.
	// so, it will create the array first and then create the slice slc.
	// the other way to create a slice is to call make() function
	// a. make type and length (length = capacity)
	slice3 := make([]int, 10)
	fmt.Println(slice3[0])
	// b. make type, length, and capacity
	slice4 := make([]int, 10, 15)
	fmt.Println(slice4[0])
	//append function
	// * size of a slice can be increased by append()
	// * adds elements to the end of a slice, even increase the size of the underlying array if necessary
	//The following line will give "panic: runtime error: index out of range [10] with length 10"
	// fmt.Println(slice4[10])
	// but after adding slice4 = append(slice4, 5), the following line will not give any error, as it will append
	// fmt.Println(slice4[10])
	slice4 = append(slice4, 5)
	fmt.Println(slice4[10])
}

func maps() {
	fmt.Println("TESTING MAPS")
	var map1 = make(map[string]int)
	map1["jane"] = 456
	fmt.Println(map1["jane"])
	map1["michel"] = 987
	map1["marie"] = 555
	for key, val := range map1 {
		fmt.Println("key: " + key + ", value: " + strconv.Itoa(val))
	}

	var map2 = map[string]int{
		"joe": 123}
	fmt.Println(map2["joe"])
	fmt.Printf("map2's length is %d'\n", len(map2))
	delete(map2, "joe")
	//two-value assignment test for exstence of the key
	id, p := map2["joe"]
	if p {
		fmt.Printf("Joe's id:%d\n", id)
	} else {
		fmt.Printf("Joe is not existing in the map!\n")
	}
}

func structs() {
	// * Aggregate data type
	// * Groups together other objects of arbitrary type
	type Person struct {
		name    string
		address string
		phone   string
	}

	p1 := new(Person)
	p1.name = "name1"
	p1.address = "address1"
	p1.phone = "phone1"
	fmt.Println(p1.name)

	p2 := Person{name: "name2", address: "address2", phone: "phone3"}
	fmt.Println(p2.address)
}

func ioutils() {
	/*
		IOUtil functions are for small files or basic file operations
	*/
	//READ
	inputByteArray, e := ioutil.ReadFile("resources/ioutil_file_in.txt")
	inputStr := string(inputByteArray)
	if e == nil {
		fmt.Printf("File content: %s\n", inputStr)
	} else {
		fmt.Println(e)
	}
	//WRITE
	outputStr := "text written to ioutil_file_out.txt"
	outputByteArray := []byte(outputStr)
	e = ioutil.WriteFile("resources/ioutil_file_out.txt", outputByteArray, 0777)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Content successfully written to file!")
	}
}

func osFileOps() {
	//READ
	file_in, err_in := os.Open("resources/os_file_in.txt")
	if err_in != nil {
		fmt.Println(err_in)
	} else {
		byteArray := make([]byte, 5)
		for {
			numberOfBytesRead, _ := file_in.Read(byteArray)
			if numberOfBytesRead != 0 {
				slice := byteArray[0:numberOfBytesRead]
				fmt.Print(string(slice))
			} else {
				break
			}
		}
		file_in.Close()
	}
	//WRITE
	file_out, err_out := os.Create("resources/os_file_out.txt")
	if err_out != nil {
		fmt.Println(err_out)
	} else {
		outputStr := "text written to os_file_out.txt"
		outputByteArray := []byte(outputStr)
		file_out.Write(outputByteArray)
		file_out.Close()
	}
}

//Declare a variable as a func
func variableAsFunction() {
	var funcVar func(int) int
	funcVar = incFn
	fmt.Printf("funcVar(1)= %d\n", funcVar(1))
	funcVar = decFn
	fmt.Printf("funcVar(1)= %d\n", funcVar(1))
}

func incFn(x int) int {
	return x + 1
}

func decFn(x int) int {
	return x - 1
}

//Function as Arguments
func functionAsArguments() {
	fmt.Printf("applyIt(incFn, 3)= %d\n", applyIt(incFn, 3))
	fmt.Printf("applyIt(decFn, 3)= %d\n", applyIt(decFn, 3))
}
func applyIt(aFunction func(int) int, val int) int {
	return aFunction(val)
}

//Anonymous Functions
func anonymousFunction() {
	v := applyIt(func(x int) int { return x + 1 }, 2)
	fmt.Printf("v:= applyIt(func (x int) int {return x + 1}, 2)= %d\n", v)
}

//Returning Function: Function Defines a Function
func testFunctionReturnsFunction() {
	Dist1 := MakeDistOrigin(0, 0)
	Dist2 := MakeDistOrigin(2, 2)
	fmt.Println(Dist1(2, 2))
	fmt.Println(Dist2(2, 2))
}

//Closure: Function + its environment (scope)
//When functions are passed/returned, their environment comes with them
//origin_x and origin_y are in the closure of fn()
func MakeDistOrigin(origin_x, origin_y float64) func(float64, float64) float64 {
	function := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-origin_x, 2) + math.Pow(y-origin_y, 2))
	}
	return function
}

//variadic: variable argument number
//* functions can take a variable number of arguments
//* uses ellipses ... to specify
//treated as a slice inside function
func testVariadic() {
	fmt.Printf("getMax(1,3,5,4): %d\n", getMax(1, 3, 5, 4))
	slice := []int{1, 3, 5, 3}
	fmt.Printf("getMax(slice): %d\n", getMax(slice...))
}
func getMax(values ...int) int {
	maxValue := -1
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

//Deferred Function Calls
// * call can be deferred until the surrounding function completes
// * typically used for cleanup activities
// * arguments of a deferred call are evaluated immediately
func testDeferred() {
	i := 1
	defer fmt.Println(i + 1)
	i++
	fmt.Println("Hello!")
}

//INTERFACES
//Terminology: in the following example, the dynamic-type is Dog and dynamic-value is d1
//interface with Nil dynamic value (dynamic-type is Dog)
//var s1 Speaker
//var d1 *Dog
//s1=d1
func testInterfaces() {
	fmt.Println("TESTING INTERFACES")
	var s1 Speaker
	var d1 = Dog1{"Brian"}
	s1 = d1
	s1.Speak()
	typeAssertionViaIfElse(s1)
	typeAssertionViaSwitch(s1)
	PrintMe(d1)
}

type Speaker interface{ Speak() }
type Dog1 struct{ name string }
type Cat1 struct{ name string }

func (d Dog1) Speak() {
	fmt.Println(d.name)
}
func (d Cat1) Speak() {
	fmt.Println(d.name)
}

//Type Assertion
func typeAssertionViaIfElse(s Speaker) {
	dog, ok := s.(Dog1)
	if ok {
		fmt.Printf("Type-Assertion(via-if-else): Speaker is a dog-> %v\n", dog)
	}
	cat, ok := s.(Cat1)
	if ok {
		fmt.Printf("Type-Assertion(via-if-else): Speaker is a car-> %v\n", cat)
	}
}
func typeAssertionViaSwitch(s Speaker) {
	switch sp := s.(type) {
	case Dog1:
		fmt.Printf("Type-Assertion(via-switch): Speaker is a dog-> %v\n", sp)
	case Cat1:
		fmt.Printf("Type-Assertion(via-switch): Speaker is a car-> %v\n", sp)
	}
}

//Empty Interface:
// * empty interface specifies no methods
// * all types satisfy the empty interface
// * use it to have a function accept any type as a parameter
func PrintMe(val interface{}) {
	fmt.Printf("Printing from function with Empty-Interface param: %v\n", val)
}

//Concurrency
// - Concurrent: start and end times overlap
//   * Concurrent tasks may be executed on the same hardware
// - Parallel: execute at exactly the same time
//   * Parallel tasks must be executed on different hardware
//Concurrent Programming
// - Programmer determines which tasks can be executed in parallel
// - Tasks mapped to hardware by
//   * Operating System
//   * Go run-time scheduler
func testWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	wg.Wait()
	fmt.Println("Main Routine")
}
func foo(wg *sync.WaitGroup) {
	fmt.Println("New Routine")
	wg.Done()
}

//TODO: add examples for (last 2w)
// 1. channel (buffered, unbuffered, multiple channel read, selective channel
// 2. mutual exclusion (mutex)
// 3. Synchronization Initialisation Sync.Once
// 4. Deadlock
