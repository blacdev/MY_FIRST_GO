<<<<<<< HEAD
// // This how we specify what this code is to do. Either to run or it is a function
// package main
//
// // this is how to import a single package
// import "fmt"
//
// //This is how to import multiple packages and also assign it to a variable
// // import (
// //     print "fmt"
// //
// // )
//
//
// /* This is how to write a function in Golang
//     It also explains that any file that has the package "main" imported is
//     is goin to run the file when the go run *filename* is typed in the CLI
// */
//
// func main(){
//     var x int = 1
//     var y int
//     var ip *int          // Ip is currently set to point to any address assigned to it
//     ip = &x              // Ip is assigned the address of integer x
//     fmt.Println(ip)      // shows the address integer x is assigned to
//     y = *ip              // variable y is now assigned the integer value located in the address assigned to ip
//     fmt.Printf("%d\n", y)
//
//     ptr:= new(int)       // new() creates a variable and returns a pointer to the variable
//     y = *ptr
//     fmt.Println(y)
//     *ptr = 3
//     fmt.Println(*ptr)
//     g()
//     f()
//
// }
//
// func g(){
//     var x int = 4
//     fmt.Println(x)
//
// }
// /*
// Printing
// */
//
// func f(){
//     fmt.Printf("Hi")
//     x:="Joe"
//     fmt.Printf("Hi %s",  x + "\n")
// }
=======
package main


import(
    "fmt"
    "rsc.io/quote"
)

func main(){
    var x int = 5
    var y int = 7
    v := 4
    h :=5
    s := v + h
    fmt.Println(s)
    var sum int = x + y
    a:= [5]int
    fmt.Println(a)
    if x > 6{

      fmt.Println("Greater than6")
    }else if x < 2{
      fmt.Println("we move")
    }else{

      fmt.Println("whatever")
    }

    fmt.Println(sum)
    fmt.Println("Hello, World")
    fmt.Println(quote.Go())
}
>>>>>>> c71d7ee8f68abc91be8830808161134fd147e25b
