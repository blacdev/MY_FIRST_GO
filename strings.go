package main

import( F "fmt"
        S "strings"
      )


func main()  {
  // Assign string to variable
  var greetings = "Hello world!"
  greet := []string{"Hello", "world!"}

  // Print string
  F.Printf("normal string: %s\n", greetings)

  // print string in hex values
  F.Printf("hex bytes: ")

  // loop to convert each  character to a byte
  for i:=0; i < len(greetings); i++ {
    F.Printf("%x ", greetings[i])
  }
  F.Printf("\n")

  // print length of string
  F.Println(len(greetings))

  // join two different words together using builtin function join
  F.Println(S.Join(greet, " "))
}
