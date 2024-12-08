package main

import(
  "fmt"
)

func main(){
  dynamic_arr := make([]int, 5)
  fmt.Println(dynamic_arr)
  dynamic_arr[0] = 1
  dynamic_arr[1] = 3
  dynamic_arr[2] = 32
  dynamic_arr[3] = 5
  dynamic_arr[4] = 2
  fmt.Println(dynamic_arr)
  dynamic_arr = append(dynamic_arr, 34,324,32445)
  fmt.Println(dynamic_arr)
}
