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

  if (a == 2){
    F.Printf("True %d a\n", a)
  }
  if (b == 1){
    F.Printf("True %d b\n", b)
  }
  if (a != b){
    F.Printf("False\n")
  }

// We printThe answers for aech operation carried out
  F.Printf("%d is the final answer when you add up a an b together.\n", c)
  F.Printf("%f is the final answer when you divide a and b together.\n", d)
  F.Printf("%d is the final answer when you subtract a and b together.\n", f)
  F.Printf("%d is the final answer you get when you multiply a and b together.\n", e)
  F.Printf("%d\n", g)
  F.Printf("%d\n", a)
  F.Printf("%d\n", b)
}
