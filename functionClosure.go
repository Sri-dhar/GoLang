package main

import(
  "fmt"
)

func doCompute(fn func(int,int) int) int{
  return fn(3,5)
}

func main(){
  add := func(x,y int) int{
    return x+y
  }
  fmt.Println(doCompute(add))

  //using closure
  x := 0
  increment := func() int{
    x++
    return x
  }
  for i:=0; i<4; i++{
    fmt.Println(increment())
  }
  fmt.Println("After incrementing 4 times: ")

}


