package main

import "fmt"

func main2() {
	fmt.Println("Hello, World!")

}

/*
is this a comment?

func main()
{
	fmt.Println("Hello, World!")
}

this gives error

*/

func main() {
	fmt.Println("Hello, World!")
	var a int = 43
	fmt.Println("Go " + string(a) + "Language")

	//print int and string in sinngle line
	fmt.Println("Go", a, "Language")
	var aa float32 = 43.2
	fmt.Println("No. of variants of a mutating bacteria is", aa)
}
