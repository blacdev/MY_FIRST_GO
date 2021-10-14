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
