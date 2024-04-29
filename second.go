package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	//all data types in go
	var a int = 10
	var b float64 = 3.14
	var c string = "Hello"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//multiple variable declaratio
	var x, y int = 10, 20
	fmt.Println(x)
	fmt.Println(y)

	//take input of variables
	var input int
	fmt.Printf("\nEnter a number:\n")
	fmt.Scanln(&input)
	fmt.Println(input)
}
