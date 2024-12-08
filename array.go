package main

import(
  "fmt"
)

func main(){
  var a[5] int = [5] int{1,3,32,5,2}
  fmt.Println(a)
  fmt.Println(a[2])

  //array operations
  var sum int = 0
  for i:=0; i<5; i++{
    sum += a[i]
  }

  fmt.Println("Sum: ", sum)

  //array length
  fmt.Println("Length: ", len(a))

  //array of strings
  var b[2] string
  b[0] = "Hello"
  b[1] = "World"
  fmt.Println(b)

}
