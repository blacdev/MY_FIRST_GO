package main

import  F "fmt"

func main(){
  var a = 10
  var ip *int // declare ip as pointer variable
  var fp *float32 // declare fp as pointer variable
  F.Printf("address of a is %x\n", &a)

  // move the value at a to ip
  ip = &a
  
  F.Printf("The value at poisition a is now at position ip %d\n", *ip)
  //checking the type of variable ip is
  F.Printf("%T\n", ip)
  // print the value initialized with fp
  F.Printf("%x\n",fp)

}
