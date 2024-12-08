package main

import "fmt"

func Sqrt(x float64) float64{
  var z = 1.0
  for{
    z = z - (z*z - x)/(2*z)
    fmt.Println("the value of z is : ",z," and the value of z*z - x is : ",z*z - x)
    if z*z - x < 0.000000001{
      break
    }
  }
  return z
}

func main(){
  fmt.Println(Sqrt(2))
}
