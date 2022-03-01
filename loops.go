// package main
//
// import F "fmt"
//
// func main(){
//   //create variables
//   var b int = 15
//   var a, j int
//   numbers := [6]int{1, 2, 3, 5}
//
//   //  loop through a till is one less than 9
//   for a := 0; a < 10; a++{
//     F.Printf("The value of a is: %d\n", a)
//   }
//
//   F.Printf("The end of first loop\n")
//
//   //  loop through a till is one less than b
//   for a < b{
//     a++
//     F.Printf("The value of a is: %d\n", a)
//   }
//
//   F.Printf("The end of the second loop \n")
//
//   // This is a key value pair situtation. I do not fully understand it yet
//   for i,s := range numbers{
//     F.Printf("This is the vaue of x = %d at position %d\n", s, i)
//   }
//   F.Printf("The end of the third loop\n")
//
//   // a loops through 1 to 99 while j is looking for the value that matches a
//   for a = 0; a < 100; a ++{
//     for j = 0; j <= a; j++{
//       F.Printf("a = %d and j = %d\n", a, j)
//     }
//   }
//   for a = 0; a < 100; a ++{
//
//     if a <= 50{
//       F.Printf("The value of a is %d\n", a)
//       continue
//       if a > 40{
//         F.Printf("The value of a is %d\n", a)
//         break
//       }
//     }
//
//   }
//   F.Printf("Hello \n")
// }



// package main
//
// import F "fmt"
//
// func main(){
//    a := 10
//
//    LOOP: for a < 20{
//      if  a == 15{
//        a+= 1
//        goto LOOP
//      }
//      F.Printf("value of a: %d\n", a)
//      a++
//    }
// }

package main

import (F "fmt"
        com "os"
)

func main()  {
  var value string
  for {
    F.Println("Please enter your name: ")
    F.Scanln(&value)

    if value == "seye"{
      F.Println("Hello " + value)
      com.Exit(3)
    }
  }
}
