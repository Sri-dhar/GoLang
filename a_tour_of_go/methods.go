package main

import "fmt"
import "math"

type student struct{
  name string
  roll int
}

func (s student) display() bool{
  fmt.Println("Name of the student: ", s.name)
  fmt.Println("Roll number of the student: ", s.roll)
  return true
}

func (s student) area() float64{
  return math.Pi * float64(s.roll) * float64(s.roll)
}

func main(){
  s := student{"John", 12}
  s.display()
  fmt.Println("Area of the circle: ", s.area())
}
