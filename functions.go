package main

import (F "fmt"
         "math"
      )

// Function to by vlaue
func max(num1, num2 int) (int) {
  var result int

  result = num1
  num1 = num2
  num2 = result

  return num2
}

// Function to swap value by refrence using their address location in memory
func min(num1 *int, num2 *int) {
  var result int

  result = *num1
  *num1 = *num2
  *num2 = result


}

func main(){
  a := 100
  b := 300

  // learning to call by value
  F.Printf("Before swap, value of a: %d\n", a)
  F.Printf("Before swap, value of b: %d\n", b)
  max(a, b)
  F.Printf("After swap, value of a: %d\n", a)
  F.Printf("After swap, value of b: %d\n", b)
  // F.Printf("The max value is: %d\n", ret)

  F.Println("\n")
  F.Println("\n")

  // learning to call by refrence
  F.Printf("Before swap2, value of a: %d\n", a)
  F.Printf("Before swap2, value of b: %d\n", b)
  min(&a, &b)
  F.Printf("After swap2, value of a: %d\n", a)
  F.Printf("After swap, value of b: %d\n", b)

  // create a function as a variable
  getSquareRoot := func(x float64) float64{
    return math.Sqrt(x)
  }
  F.Println(getSquareRoot(9))

  // passing data characteristics of a cricle
  circle := Circle{x:0, y:0, radius: 5}

  // calling the method of a circle that has been created
  F.Printf("Area of a crcle is %f\n", circle.area())
}

// definig the characteristics of a circle
type Circle struct{
  x, y, radius float64
}

// define a method that makes use of the characteristics of a Circle
func (circle Circle) area() float64  {
  return math.Pi * circle.radius * circle.radius
}



// I learnt:
// -Functions and how to they can be used:
// - their call type either by value or by reference
