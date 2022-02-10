package main

import F "fmt"






func main() {
  var a, b, c, e, f, g  int
  var d float32

  F.Println("Enter a value for position 'a': ")
  F.Scanln(&a)

  F.Println("Enter a value for position 'b':")
  F.Scanln(&b)

  c = a + b // gives the sum of a and b
  d = float32(b) / float32(a) // give the division of a and b

  f = a - b // gives the subtraction of a and b

  e = a * b // give the multiplaction of a and b
  g = b % a // give the remainder after integer division
  a++ // increments a by one
  b-- // decrements b by one



// We printThe answers for aech operation carried out
  F.Printf("%d is the final answer when you add up a an b together.\n", c)
  F.Printf("%f is the final answer when you divide a and b together.\n", d)
  F.Printf("%d is the final answer when you subtract a and b together.\n", f)
  F.Printf("%d is the final answer you get when you multiply a and b together.\n", e)
  F.Printf("%d\n", g)
  F.Printf("%d\n", a)
  F.Printf("%d\n", b)

// if statements
  //  check if the value of a is 2
  if (a == 2){

    // print result
    F.Printf("True %d a\n", a)
  }

  // check if the value of b is 1
  if (b == 1){

    // print result
    F.Printf("True %d b\n", b)
    b = 2
  }

// if else statement
  // Check if a is not equal to b
  if (a != b){
    // print result
    F.Printf("False\n")
  }    else    {
    F.Printf("True")
  }


}
