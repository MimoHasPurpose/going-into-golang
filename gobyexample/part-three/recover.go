// to recover from panic use recover 
package main
import "fmt"
func mayPanic(){
	panic("a porblem")
}


func main(){
	defer func(){
		if r:= recover(); r!=nil{
			fmt.Println("Recovered. Error:\n",r)
		}

	}()
	mayPanic()
	fmt.Println("After panic()")
}