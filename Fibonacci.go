/*
Writing fibonacci function in Golang
-Given 2 intial values 0 and 1 and a limit at which the the sequnce is to print for
- the function is to add the the first and second together
- then assign the second value to position one and the vew value gotten to position
*/
package main

import (F "fmt"
       com "os"
     )

func fibonaci(amount, num1, num2 int) int{
  //
  var count int = 0
  if amount <= 0{
    com.Exit(3)
  }else if amount == 1{
    F.Println("The fibonaci sequence is ", amount)
  }else{
    for count < amount{
      F.Println(num1)
      num3 := num1 + num2
      num1 = num2
      num2 = num3
      count += 1
    }
  }
  return num1
}

func main()  {
  var s int
  F.Println("Please enter a number of how long you want the Fibonacci to be genearated:")
  F.Scanln(&s)
  F.Println("The Fibonacci sequence requested is:")
  fibonaci(s,0,1)
}
