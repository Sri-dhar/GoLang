package main

import "fmt"

func main() {
	fmt.Println("Helo world")
	var a int = 234
	fmt.Println("the value of a is : ", a)
	//take input
	var b int
	fmt.Println("Enter the value of b : ")
	fmt.Scanln(&b)

	fmt.Println("The value of b is : ", b)

	//create an array
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println("The array is : ", arr)

	//create an array with values
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("The array is : ", arr1)

	//create an array with values
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("The array is : ", arr2)

	//create an array with values

	arr3 := [5]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}

	fmt.Println("The array is : ", arr3)

	//demostrate the use of slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("The slice is : ", slice)

	//show all arrat functions
	fmt.Println("The length of the array is : ", len(arr))
	fmt.Println("The capacity of the array is : ", cap(arr))

	//do go have functions like push and pop
	slice = append(slice, 6)
	fmt.Println("The slice is : ", slice)

	//create a map

	mp := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println("The map is : ", mp)

	//create a map

	mp1 := map[string]int{}
	mp1["one"] = 1
	mp1["two"] = 2
	mp1["three"] = 3

	fmt.Println("The map is : ", mp1)

}
