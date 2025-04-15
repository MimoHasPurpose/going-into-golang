//  go supports anonymous function which can form closures.
// anonymous function are useful when u want to define a function without having to name it.

package main
import "fmt"

func intSeq() func() int{
    i:=0
    return func() int {
        i++
        return i
    }
}

func main(){
    // we call intSeq(): assigning a result (function) to nextInt, it captures its own i value, and value gets updated everytime we call nextInt.

    nextInt:=intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts:=intSeq()
    fmt.Println(newInts())
    fmt.Println(newInts())
    fmt.Println(newInts())
    fmt.Println(newInts())
    fmt.Println(newInts())
    fmt.Println(newInts())
    fmt.Println(newInts())
}