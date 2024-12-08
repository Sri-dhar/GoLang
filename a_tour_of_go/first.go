package main
import (
  "fmt"
  "time"
  "math/rand"
)

var string1 string = "A string declared Globally";

func doSomething(x int, y int) int{
  fmt.Println("Doing something")
  return x+y
}

func newFunction(x,y int, z string) (a int,b string){
  fmt.Println("doing something in the new function")
  a = x+y
  b = z+z
  return
}

func main(){
  fmt.Println("Hello World")
  fmt.Println("The time is ",time.Now())
  fmt.Println("A random number is ", rand.Intn(10))
  int1 := 4
  int2 := 5
  fmt.Println("The sum of ",int1," and ",int2," is ",doSomething(int1,int2))
  fmt.Println("callling a new function\n\\n")

  a,b := newFunction(4,5,"hello")
  fmt.Println("The sum of 4 and 5 is ",a)
  fmt.Println("The string is ",b)
  fmt.Println("The global string is ",string1)  
  
  var boolean bool = true
  fmt.Println("The boolean value is ",boolean)

  var float1 float32 = 3.14
  fmt.Println("The float value is ",float1)

  var int3 int = 5
  fmt.Println("The integer value is ",int3)

  var complex1 complex64 = 3+4i
  fmt.Println("The complex value is ",complex1)
}
