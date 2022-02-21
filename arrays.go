package main

import( F "fmt")

func main(){
  // Declare array variables
  balance := []float64 {12.3, 4.5, 20.0, 45.7}
  // Dclare variable used to store sum of array
  var value float64 = 0

  // loop through array and sum it up
  for i:=0; i < len(balance); i++{

    // Add value in position i to value and assign vvalue to value
    value += balance[i]

    //print each variable in the array
    F.Println(balance[i])
  }
  // Print total
  F.Println(value)
}
