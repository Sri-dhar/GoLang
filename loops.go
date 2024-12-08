package main

import(
  "fmt"
)

func main(){
  for i:= 0; i<10; i++{
    fmt.Println(i)
  }
  var sum = 0
  for i:= 0; i<10; i++{
    sum += i
  }
  fmt.Println("The sum is ",sum)
}
