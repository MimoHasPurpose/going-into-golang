
package main
import (
    "fmt"
    "math/rand/v2"
)

func main(){

	fmt.Print(rand.IntN(100)," ")
	fmt.Print(rand.IntN(100)," ")
	fmt.Print(rand.Float64()," ")
	s2:=rand.NewPCG(42,1024)
	r2:=rand.New(s2)
	fmt.Print(r2.IntN(100),",")
	fmt.Print(r2.IntN(100),",")
	s3:=rand.NewPCG(42,1025)
	r3:=rand.New(s3)
	fmt.Print(r3.IntN(100),",")
}

