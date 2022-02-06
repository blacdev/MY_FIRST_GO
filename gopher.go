// specify type of code
package main

// Import  formatted I/O function
import F "fmt"

// Beginning of main code
func main(){
  const s = 10
  t := 23
  var u = s + t
  var i, j int
  F.Println("Enter a number for i: ")
  F.Scanln(&i)
  F.Println("Enter a number for j: ")
  F.Scanln(&j)
  F.Println(i * j)
  F.Printf("s is of type %T\n", s)
  F.Printf("value of u is %d\n", u)
  F.Printf("Hello, World!\a\n")
  // F.Println("    `.-::::::-.`   ")
  // F.Println(".:-::::::::::::::-:.")
  // F.Println("`_:::    ::    :::_`")
  // F.Println(" .:( ^   :: ^   ):.")
  // F.Println(" `:::   (..)   :::.")
  // F.Println(" `:::::::UU:::::::`")
  // F.Println(" .::::::::::::::::.")
  // F.Println(" O::::::::::::::::O")
  // F.Println(" -::::::::::::::::-")
  // F.Println(" `::::::::::::::::`")
  // F.Println("  .::::::::::::::.")
  // F.Println("    oO:::::::Oo")

}
