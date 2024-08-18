package main

import "fmt"
import "unicode/utf8"

func main() {
	fmt.Println("Hello World!")
	var intNum int16 = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello \nWorld"\
	fmt.Println(myString)

	fmt.Println(len("A"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = true
	fmt.Println(myBoolean)

	var intNum3 rune
	fmt.Println(intNum3)

	var myVar string = foo()
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myconst string = "const value"
	fmt.Println(myconst)

func printMe(printValue string){
	fmt.Println("Hello World")
}
}
