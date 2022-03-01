// package main
//
// import( F "fmt")
//
// func main(){
//   // Declare array variables
//   balance := []float64 {12.3, 4.5, 20.0, 45.7}
//   // Dclare variable used to store sum of array
//   var value float64 = 0
//
//   // loop through array and sum it up
//   for i:=0; i < len(balance); i++{
//
//     // Add value in position i to value and assign vvalue to value
//     value += balance[i]
//
//     //print each variable in the array
//     F.Println(balance[i])
//   }
//   // Print total
//   F.Println(value)
// }

/*
Revisit to arrays
*/
package main

import( "fmt"
        com "os"
)
func main() {
  // intialize a variable that tells us the size of the of the array
   var size int

   // Requesting for the the size
   fmt.Println("Please enter the amount of fibonacci to be printed:")
   fmt.Scanln(&size)

   //Pass size into fibonaciIntoArray function
   avg := fibonaciIntoArray(size, 0, 1)

   // Print array
   fmt.Println( "fibonacci in array: ");
   fmt.Println(avg)

   // Pass array into printArray function
   fmt.Println( "Array listing is: ");
   printArray(avg, size)

   /* output the returned value */

}



func fibonaciIntoArray(size, num1, num2 int) []int {
  arr := make([]int, size)
  count := 0
  if size == 0{
    fmt.Println("Invalid entry")
    com.Exit(3)
  }else if size == 1{
    fmt.Println("the fibonacci sequence is:")
    fmt.Println(num1)
  }else {
    fmt.Println("The Fibonacci sequence is:")
    for count < size{
      // fmt.Println(num1)
      num3 := num1 + num2
      arr[count] = num1
      num1 = num2
      num2 = num3
      count += 1
      }
  }
  return arr
}



func printArray(arr []int, size int) int {
  var i int
  for i := 0; i < size; i++ {
    fmt.Println(arr[i])
  }
  return arr[i]
}
