package main

import F "fmt"

func main(){
  var a int = 20
  b := 10
  var t int
  var grade string
  var marks int
  var x interface{}


  F.Println("Enter a value of T: ")
  F.Scanln(&t)

// if else if statements
  if (t == a){

    F.Printf("%d is equal a\n", t)
  } else if (t == b){
    F.Printf("%d is equal to b\n", t)
  } else {
    F.Printf("None of the values is matching %d\n", t)

  }

// nested if statements
  if (t == 60){
    F.Println("We are about to do something magical\n")
   if (t % 2 == 0){
    F.Println("We just got a value of 0")
  }
 }
 F.Printf("Value of t is %d\n", t)
 F.Println("Enter a value of grade: ")
 F.Scanln(&marks)
 switch marks {
 case 90, 100:grade = "A"
 case 80:  grade = "B"
 case 50, 60, 70: grade = "C"
 default: grade = "D"

 }
 switch{
 case grade == "A", grade == "a" :
   F.Printf("Excellent!\n")
 case grade == "B", grade == "b" :
   F.Printf("Well done!\n")
 case grade == "C", grade == "c" :
   F.Printf("Well done!\n")
 case grade == "d", grade == "D":
   F.Printf("You passed\n")
 case grade == "F", grade == "f":
   F.Printf("you need to take the paper again\n")
 default:
   F.Printf("Invalid grade\n");
  }
  F.Printf("Your grade is %s\n", grade)
  switch i:= x.(type) {
  case nil:
    F.Printf("type of x : %T", i)
  case int:
    F.Printf("x is int")
  case float32:
    F.Printf("x is float 32")
  case float64:
    F.Printf("x is float 64")
  case func(int) float64:
    F.Printf("x is float 32")
  case bool, string:
    F.Printf("x is bool or string")
  default:
    F.Printf("don't know the type")

  }
}
