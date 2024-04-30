package main

import "fmt"

func starPrinter(a int){
	if(a==0){
		return
	}
	for i:= range(a){
		i++
		fmt.Print("*")
	}

	fmt.Println()
	starPrinter(a-1)
}

func main(){
	var a int
	fmt.Println("Enter the number of stars you want to print")
	fmt.Scan(&a)
	starPrinter(a)
}